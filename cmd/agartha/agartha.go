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
	"github.com/murilocosta/agartha/internal/application/auth"
	"github.com/murilocosta/agartha/internal/application/reports"
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
		configFilePath := basepath + "/../../configs/config.yml"

		if len(os.Args) > 2 {
			configFilePath = os.Args[2]
		}

		err = core.LoadConfig(configFilePath, &config)
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

	// If cache is enabled, create an instance of Redis client
	if config.Cache.Enabled {
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
}

func main() {
	// Initialize database repositories
	survRepo := persistence.NewPostgresSurvivorRepository(dbCli)
	itemRepo := persistence.NewPostgresItemRepository(dbCli)
	infectRepo := persistence.NewPostgresInfectionRepository(dbCli)
	invRepo := persistence.NewPostgresInventoryRepository(dbCli)
	tradeRepo := persistence.NewPostgresTradeRepository(dbCli)
	credRepo := persistence.NewPostgresCredentialsRepository(dbCli)
	reportRepo := persistence.NewPostgresReportRepository(dbCli)

	// Initialize domain services
	invServ := domain.NewInventoryService(invRepo)

	// Initialize application use cases
	regSurUC := application.NewRegisterSurvivor(survRepo, itemRepo)
	updLocUC := application.NewUpdateLastLocation(survRepo)
	ftcSurvDtlUC := application.NewFetchSurvivorDetails(survRepo)
	ftcSurvLstUC := application.NewFetchSurvivorList(survRepo)
	ftcSurvInvUC := application.NewFetchSurvivorInventory(invRepo)
	ftcAvlbItms := application.NewFetchAvailableItems(itemRepo)
	flgInfUC := application.NewFlagInfectedSurvivor(survRepo, infectRepo, invServ)
	trdUC := application.NewTradeItems(survRepo, tradeRepo)
	trdAccUC := application.NewTradeItemsAccept(tradeRepo)
	trdRejUC := application.NewTradeItemsReject(tradeRepo)
	trdCancUC := application.NewTradeItemsCancel(tradeRepo)
	trdHstUC := application.NewFetchSurvivorTradeHistory(tradeRepo, itemRepo)
	survSgn := auth.NewSignUpSurvivor(credRepo, itemRepo)
	survLgn := auth.NewLoginSurvivor(credRepo)
	showInfectPerc := reports.NewShowInfectedPercentage(reportRepo)
	showNonInfectPerc := reports.NewShowNonInfectedPercentage(reportRepo)

	// Initialize request handlers
	registerSurvivor := transport.NewRegisterSurvivorCtrl(regSurUC)
	updateLastLocation := transport.NewUpdateLastLocationCtrl(updLocUC)
	fetchSurvivorProfile := transport.NewFetchSurvivorProfileCtrl(ftcSurvDtlUC)
	fetchSurvivorDetails := transport.NewFetchSurvivorDetailsCtrl(ftcSurvDtlUC)
	fetchSurvivorList := transport.NewFetchSurvivorListCtrl(ftcSurvLstUC)
	fetchSurvivorInventory := transport.NewFetchSurvivorInventoryCtrl(ftcSurvInvUC)
	fetchAvailableItems := transport.NewFetchAvailableItemsCtrl(ftcAvlbItms)
	flagInfectedSurvivor := transport.NewFlagInfectedSurvivorCtrl(flgInfUC)
	tradeItems := transport.NewTradeItemsCtrl(trdUC)
	tradeItemsAccept := transport.NewTradeItemsAcceptCtrl(trdAccUC)
	tradeItemsReject := transport.NewTradeItemsRejectCtrl(trdRejUC)
	tradeItemsCancel := transport.NewTradeItemsCancelCtrl(trdCancUC)
	tradeItemsHistory := transport.NewFetchSurvivorTradeHistoryCtrl(trdHstUC)
	showInfectedPercentage := transport.NewShowInfectedPercentageCtrl(showInfectPerc)
	showNonInfectedPercentage := transport.NewShowNonInfectedPercentageCtrl(showNonInfectPerc)

	// Initialize auth handlers
	survivorSignUp := transport.NewSurvivorSignUpCtrl(survSgn)
	survivorLogin := transport.NewSurvivorLoginCtrl(survLgn)
	checkSurvivorPermission := transport.NewCheckSurvivorPermissionCtrl()

	// Register the controllers
	handlersConfig := infrastructure.NewServerHandlersConfig()
	handlersConfig.Post("/api/register", survivorSignUp)
	handlersConfig.Get("/api/items", fetchAvailableItems)
	handlersConfig.PostProtected("/api/survivors", registerSurvivor)
	handlersConfig.PutProtected("/api/survivors/:survivorId", updateLastLocation)
	handlersConfig.GetProtected("/api/survivors/profile", fetchSurvivorProfile)
	handlersConfig.GetProtected("/api/survivors/:survivorId", fetchSurvivorDetails)
	handlersConfig.GetProtected("/api/survivors/:survivorId/trades", tradeItemsHistory)
	handlersConfig.GetProtected("/api/survivors/:survivorId/items", fetchSurvivorInventory)
	handlersConfig.GetProtected("/api/survivors", fetchSurvivorList)
	handlersConfig.PostProtected("/api/survivors/report-infection", flagInfectedSurvivor)
	handlersConfig.PostProtected("/api/trades", tradeItems)
	handlersConfig.PostProtected("/api/trades/:tradeId/accept", tradeItemsAccept)
	handlersConfig.PostProtected("/api/trades/:tradeId/reject", tradeItemsReject)
	handlersConfig.PostProtected("/api/trades/:tradeId/cancel", tradeItemsCancel)
	handlersConfig.Get("/api/reports/infected-survivors", showInfectedPercentage)
	handlersConfig.Get("/api/reports/non-infected-survivors", showNonInfectedPercentage)

	// Create the authentication middleware
	middleware := infrastructure.NewAuthMiddleware(
		config.Auth.Realm,
		config.Auth.SecretKey,
		config.Auth.TokenTimeout,
		config.Auth.RefreshTimeout,
	)

	authHandler, err := middleware.Init(
		survivorLogin.HandlerFunc,
		checkSurvivorPermission.HandlerFunc,
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
