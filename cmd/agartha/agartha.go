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

func init() {
	var config core.Config
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
	survRepo := persistence.NewPostgresSurvivorRepository(dbCli)
	itemRepo := persistence.NewPostgresItemRepository(dbCli)
	infectRepo := persistence.NewPostgresInfectionRepository(dbCli)
	invRepo := persistence.NewPostgresInventoryRepository(dbCli)

	rsuc := application.NewRegisterSurvivor(survRepo, itemRepo)
	registerSurvivor := transport.NewRegisterSurvivorCtrl(rsuc)

	ulluc := application.NewUpdateLastLocation(survRepo)
	updateLastLocation := transport.NewUpdateLastLocationCtrl(ulluc)

	fsluc := application.NewFetchSurvivorList(survRepo)
	fetchSurvivorList := transport.NewFetchSurvivorListCtrl(fsluc)

	invServ := domain.NewInventoryService(invRepo)
	fisuc := application.NewFlagInfectedSurvivor(survRepo, infectRepo, invServ)
	flagInfectedSurvivor := transport.NewFlagInfectedSurvivorCtrl(fisuc)

	// Create an instance of the application server
	r := gin.Default()
	s := infrastructure.NewServer(r)
	s.Register(infrastructure.ServerPost, "/api/survivors", registerSurvivor)
	s.Register(infrastructure.ServerGet, "/api/survivors", fetchSurvivorList)
	s.Register(infrastructure.ServerPost, "/api/survivors/report-infection", flagInfectedSurvivor)
	s.Register(infrastructure.ServerPost, "/api/survivors/:survivorId", updateLastLocation)
	s.Run()
}
