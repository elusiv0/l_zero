package redis

import (
	"fmt"

	"github.com/elusiv0/wb_tech_l0/internal/config"
	"github.com/redis/go-redis/v9"
)

func New(cfg *config.Redis) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: "",
		DB:       0,
	})

	return client
}
