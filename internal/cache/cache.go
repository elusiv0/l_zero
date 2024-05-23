package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	orderDto "github.com/elusiv0/wb_tech_l0/internal/dto/order"
	customErrors "github.com/elusiv0/wb_tech_l0/internal/errors"
	"github.com/redis/go-redis/v9"
)

var (
	ErrNotFound = customErrors.New(http.StatusNoContent, "Order not found")
)

type Cache struct {
	redis  *redis.Client
	logger *slog.Logger
}

func New(
	redis *redis.Client,
	logger *slog.Logger,
) *Cache {
	cache := &Cache{
		redis:  redis,
		logger: logger,
	}

	return cache
}

func (cache *Cache) AddOrder(ctx context.Context, order orderDto.Order) error {
	uid := order.OrderUid
	data, err := json.Marshal(order)
	if err != nil {
		return fmt.Errorf("Cache - AddOrder - Marshall json: %w", err)
	}
	cache.logger.Info("adding order to cache", slog.Any("uid", uid))
	if err := cache.redis.Set(ctx, uid, string(data), 0).Err(); err != nil {
		return fmt.Errorf("Cache - AddOrder - redis Set - failed to add %s: %w", uid, err)
	}
	cache.logger.Info("order added to cache", slog.Any("uid", uid))
	return nil
}

func (cache *Cache) GetOrder(ctx context.Context, uid string) (string, error) {
	cache.logger.Info("searching order in cache", slog.Any("uid", uid))
	val, err := cache.redis.Get(ctx, uid).Result()
	if err != nil {
		if err == redis.Nil {
			return val, ErrNotFound
		}
		return val, fmt.Errorf("Cache - GetOrder: %w", err)
	}
	return val, nil
}
