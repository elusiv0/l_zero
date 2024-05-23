package main

import (
	"log"

	"github.com/elusiv0/wb_tech_l0/internal/app"
	"github.com/elusiv0/wb_tech_l0/internal/cache"
	"github.com/elusiv0/wb_tech_l0/internal/config"
	orderConsumer "github.com/elusiv0/wb_tech_l0/internal/consumer/order"
	orderRepository "github.com/elusiv0/wb_tech_l0/internal/repo/order"
	router "github.com/elusiv0/wb_tech_l0/internal/router"
	orderService "github.com/elusiv0/wb_tech_l0/internal/service/order"
	"github.com/elusiv0/wb_tech_l0/internal/util"
	"github.com/elusiv0/wb_tech_l0/pkg/httpserver"
	"github.com/elusiv0/wb_tech_l0/pkg/logger"
	"github.com/elusiv0/wb_tech_l0/pkg/nats"
	"github.com/elusiv0/wb_tech_l0/pkg/postgres"
	"github.com/elusiv0/wb_tech_l0/pkg/redis"
	"github.com/joho/godotenv"
)

func main() {
	//load env variables
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("error with load env variables " + err.Error())
	}

	//building config
	config, err := config.NewConfig()
	if err != nil {
		log.Fatal("error with get config " + err.Error())
	}

	//building logger
	logger := logger.New("local")

	//building postgres
	pg, err := postgres.New(
		postgres.NewConnectionConfig(
			config.Postgres.Host,
			config.Postgres.Port,
			config.Postgres.User,
			config.Postgres.Password,
			config.Postgres.Name,
		),
		logger,
		postgres.ConnAttempts(config.Postgres.ConnectionAttempts),
		postgres.ConnTimeout(config.Postgres.ConnectionTimeout),
		postgres.MaxPoolSz(config.Postgres.MaxPoolSz),
	)
	if err != nil {
		log.Fatal("error with set up pg connection " + err.Error())
	}
	//building redis
	redis := redis.New(&config.Redis)

	//building nats-client
	nats, err := nats.New(&config.Nats)
	if err != nil {
		log.Fatal("error with set up nats-client " + err.Error())
	}

	//building repo
	orderRepository := orderRepository.New(pg, logger)

	//building service
	orderService := orderService.New(orderRepository, logger)

	//building cache
	cache := cache.New(redis, logger)

	//fill cache
	if err = util.FillCacheFromDb(cache, orderRepository, logger); err != nil {
		log.Fatal(err.Error())
	}

	//building order consumer
	orderConsumer, err := orderConsumer.New(orderService, logger, nats, cache)
	if err != nil {
		log.Fatal("error with set up order consumer " + err.Error())
	}
	orderConsumer.Subscribe()

	//building router
	router := router.InitRoutes(cache, logger)

	httpserver := httpserver.New(
		router,
		httpserver.Port(config.Http.Port),
		httpserver.ReadTimeout(config.Http.ReadTimeout),
		httpserver.ShutdownTimeout(config.Http.ShutdownTimeout),
	)

	//building app
	app := app.New(httpserver, logger)

	//run app
	logger.Info("Starting app...")
	if err := app.Run(); err != nil {
		log.Fatal("error with starting app" + err.Error())
	}
}
