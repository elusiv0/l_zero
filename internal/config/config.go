package config

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type (
	Config struct {
		Http     Http
		Postgres Postgres
		Nats     Nats
		Redis    Redis
	}

	Http struct {
		Host            string        `envconfig:"HTTP_HOST" default:"localhost"`
		Port            string        `envconfig:"HTTP_PORT" default:"8080"`
		ReadTimeout     time.Duration `envconfig:"HTTP_READTIMEOUT" default:"5s"`
		WriteTimeout    time.Duration `envconfig:"HTTP_WRITETIMEOUT" default:"5s"`
		ShutdownTimeout time.Duration `envconfig:"HTTP_SHUTDOWNTIMEOUT" default:"3s"`
	}

	Postgres struct {
		MaxPoolSz          int           `envconfig:"PG_MAX_POOL_SIZE" default:"1"`
		ConnectionTimeout  time.Duration `envconfig:"PG_CONNECTION_TIMEOUT" default:"3s"`
		ConnectionAttempts int           `envconfig:"PG_CONNECTION_ATTEMPTS" default:"10"`
		Host               string        `envconfig:"PG_HOST" required:"true"`
		Port               string        `envconfig:"PG_PORT" required:"true"`
		User               string        `envconfig:"PG_USER" required:"true"`
		Password           string        `envconfig:"PG_PASSWORD" required:"true"`
		Name               string        `envconfig:"PG_NAME" required:"true"`
	}

	Nats struct {
		Host         string        `envconfig:"NATS_HOST" default:"localhost"`
		Port         string        `envconfig:"NATS_PORT" default:"4222"`
		ConnectWait  time.Duration `envconfig:"NATS_CONNECT_WAIT" default:"2s"`
		PingInterval int           `envconfig:"NATS_PING_INTERVAL" default:"5"`
		PingMaxOut   int           `envconfig:"NATS_PING_MAXOUT" default:"88"`
		ClusterId    string        `envconfig:"NATS_CLUSTER_ID" required:"true"`
		ClientId     string        `envconfig:"NATS_CLIENT_ID" required:"true"`
	}

	Redis struct {
		Host string `envconfig:"REDIS_HOST" default:"localhost"`
		Port string `envconfig:"REDIS_PORT" default:"6379"`
	}
)

func NewConfig() (*Config, error) {
	cfg := Config{}
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, fmt.Errorf("Config - NewConfig: %w", err)
	}

	return &cfg, nil
}
