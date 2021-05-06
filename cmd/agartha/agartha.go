package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	redis "github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/murilocosta/agartha/internal/infrastructure"
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
		err = core.LoadConfig("./configs/config.yml", &config)
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
	// Create an instance of the application server
	r := gin.Default()
	s := infrastructure.NewServer(r)
	s.Run()
}