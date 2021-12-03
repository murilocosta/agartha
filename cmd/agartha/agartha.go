package main

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"
	redis "github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/murilocosta/agartha/internal/application"
	"github.com/murilocosta/agartha/internal/core"
	"github.com/murilocosta/agartha/internal/domain"
	"github.com/murilocosta/agartha/internal/infrastructure"
	"github.com/murilocosta/agartha/internal/infrastructure/persistence"
	"github.com/murilocosta/agartha/internal/infrastructure/transport"
)

var (
	_, file, _, _ = runtime.Caller(0)
	basepath      = filepath.Dir(file)
)

var dbCli *gorm.DB
var redisCli *redis.Client
var config core.Config

func init() {
	var err error

	// Get configuration server address
	if configServerURL, ok := os.LookupEnv("CONFIG_SERVER"); ok {
		err = infrastructure.LoadConfigurationFromServer(configServerURL, &config)
	} else {
		err = core.LoadConfig(basepath+"/../../configs/config.yml", &config)
	}
	if err != nil {
		log.Fatal(err)
	}

	conn, err := core.ParseConnectionURL(&config)
	if err != nil {
		log.Fatal(err)
	}

	// Create an instance of database client
	dbCli, err = gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Create an instance of Redis client
	redisCli = redis.NewClient(&redis.Options{
		Addr:     config.Cache.Host,
		Password: config.Cache.Password,
		DB:       config.Cache.DatabaseSelection,
	})

	// Try to stablish a connection with Redis server
	_, err = redisCli.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Initialize database repositories
	survRepo := persistence.NewPostgresSurvivorRepository(dbCli)
	itemRepo := persistence.NewPostgresItemRepository(dbCli)
	infectRepo := persistence.NewPostgresInfectionRepository(dbCli)
	invRepo := persistence.NewPostgresInventoryRepository(dbCli)
	tradeRepo := persistence.NewPostgresTradeRepository(dbCli)

	// Initialize domain services
	invServ := domain.NewInventoryService(invRepo)

	// Initialize application use cases
	regSurUC := application.NewRegisterSurvivor(survRepo, itemRepo)
	updLocUC := application.NewUpdateLastLocation(survRepo)
	ftcSurvDtlUC := application.NewFetchSurvivorDetails(survRepo)
	ftcSurvLstUC := application.NewFetchSurvivorList(survRepo)
	ftcSurvInvUC := application.NewFetchSurvivorInventory(invRepo)
	flgInfUC := application.NewFlagInfectedSurvivor(survRepo, infectRepo, invServ)
	trdUC := application.NewTradeItems(survRepo, tradeRepo)
	trdAccUC := application.NewTradeItemsAccept(tradeRepo)
	trdRejUC := application.NewTradeItemsReject(tradeRepo)
	trdHstUC := application.NewFetchSurvivorTradeHistory(tradeRepo, itemRepo)

	// Initialize request handlers
	registerSurvivor := transport.NewRegisterSurvivorCtrl(regSurUC)
	updateLastLocation := transport.NewUpdateLastLocationCtrl(updLocUC)
	fetchSurvivorDetails := transport.NewFetchSurvivorDetailsCtrl(ftcSurvDtlUC)
	fetchSurvivorList := transport.NewFetchSurvivorListCtrl(ftcSurvLstUC)
	fetchSurvivorInventory := transport.NewFetchSurvivorInventoryCtrl(ftcSurvInvUC)
	flagInfectedSurvivor := transport.NewFlagInfectedSurvivorCtrl(flgInfUC)
	tradeItems := transport.NewTradeItemsCtrl(trdUC)
	tradeItemsAccept := transport.NewTradeItemsAcceptCtrl(trdAccUC)
	tradeItemsReject := transport.NewTradeItemsRejectCtrl(trdRejUC)
	tradeItemsHistory := transport.NewFetchSurvivorTradeHistoryCtrl(trdHstUC)

	// Initialize auth handlers
	signupSurvivor := transport.NewSurvivorSignUpCtrl()
	survivorLogin := transport.NewSurvivorLoginCtrl()
	checkSurvivorPermission := transport.NewCheckSurvivorPermissionCtrl()

	// Register the controllers
	handlersConfig := infrastructure.NewServerHandlersConfig()
	handlersConfig.AddHandler(infrastructure.ServerPost, "/api/register", signupSurvivor)
	handlersConfig.AddSecuredHandler(infrastructure.ServerPost, "/api/survivors", registerSurvivor)
	handlersConfig.AddSecuredHandler(infrastructure.ServerPost, "/api/survivors/:survivorId", updateLastLocation)
	handlersConfig.AddSecuredHandler(infrastructure.ServerGet, "/api/survivors/:survivorId", fetchSurvivorDetails)
	handlersConfig.AddSecuredHandler(infrastructure.ServerGet, "/api/survivors/:survivorId/trades", tradeItemsHistory)
	handlersConfig.AddSecuredHandler(infrastructure.ServerGet, "/api/survivors/:survivorId/items", fetchSurvivorInventory)
	handlersConfig.AddSecuredHandler(infrastructure.ServerGet, "/api/survivors", fetchSurvivorList)
	handlersConfig.AddSecuredHandler(infrastructure.ServerPost, "/api/survivors/report-infection", flagInfectedSurvivor)
	handlersConfig.AddSecuredHandler(infrastructure.ServerPost, "/api/trades", tradeItems)
	handlersConfig.AddSecuredHandler(infrastructure.ServerPost, "/api/trades/:tradeId/accept", tradeItemsAccept)
	handlersConfig.AddSecuredHandler(infrastructure.ServerPost, "/api/trades/:tradeId/reject", tradeItemsReject)

	// Create the authentication middleware
	middleware := infrastructure.NewAuthMiddleware(
		config.Auth.Realm,
		config.Auth.SecretKey,
		config.Auth.TokenTimeout,
		config.Auth.RefreshTimeout,
	)
	authHandler, err := middleware.Init(
		transport.CreateIdentity,
		survivorLogin.HandlerFunc,
		checkSurvivorPermission.HandlerFunc,
		transport.FormatUnauthorizedResponse,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Create an instance of the application server
	r := gin.Default()
	s := infrastructure.NewServer(r)
	s.ApplyCORS()

	// Register the authentication routes
	s.RegisterAuthHandlers(authHandler)
	s.RegisterCtrlHandlers(handlersConfig)
	s.Run()
}
