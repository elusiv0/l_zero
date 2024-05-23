package order

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/elusiv0/wb_tech_l0/internal/cache"
	"github.com/gin-gonic/gin"
)

type OrderRouter struct {
	cache  *cache.Cache
	logger *slog.Logger
}

func NewRouter(
	group *gin.RouterGroup,
	cache *cache.Cache,
	logger *slog.Logger,
) {
	orderRouter := &OrderRouter{
		cache:  cache,
		logger: logger,
	}
	group.GET("/:uid", orderRouter.getByUid)
}

func (orderRouter *OrderRouter) getByUid(c *gin.Context) {
	uid := c.Param("uid")
	resp, err := orderRouter.cache.GetOrder(c.Request.Context(), uid)
	if err != nil {
		c.Error(err)
		err = fmt.Errorf("OrderRouter - getByUid: %w", err)
		orderRouter.logger.Warn(err.Error())
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(resp))
}
