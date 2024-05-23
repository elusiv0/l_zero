package util

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/elusiv0/wb_tech_l0/internal/cache"
	"github.com/elusiv0/wb_tech_l0/internal/repo"
)

func FillCacheFromDb(cache *cache.Cache, repo repo.OrderRepo, logger *slog.Logger) error {
	ctx := context.Background()
	logger.Info("getting orders from db")
	orders, err := repo.GetAll(ctx)
	if err != nil {
		return fmt.Errorf("Util - FillCacheFromDb - failed while getting orders from db: %w", err)
	}
	for _, val := range orders {
		logger.Info("adding order to cache", slog.Any("uid", val.OrderUid))
		go func() {
			if err := cache.AddOrder(ctx, val); err != nil {
				logger.Error("Util - FillCacheFromDb " + err.Error())
			}
		}()
	}

	return nil
}
