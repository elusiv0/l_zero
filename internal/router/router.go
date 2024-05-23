package route

import (
	"log/slog"
	"net/http"

	"github.com/elusiv0/wb_tech_l0/internal/cache"
	errorsMiddleware "github.com/elusiv0/wb_tech_l0/internal/errors"
	orderRouter "github.com/elusiv0/wb_tech_l0/internal/router/order"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
)

func InitRoutes(
	cache *cache.Cache,
	logger *slog.Logger,
) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(sloggin.New(logger))
	router.Use(errorsMiddleware.ErrorsMiddleware())
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ping": "pong",
		})
	})

	orderGroup := router.Group("/api/v1/orders/")
	{
		orderRouter.NewRouter(orderGroup, cache, logger)
	}

	return router
}
