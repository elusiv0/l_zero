package repo

import (
	"context"

	"github.com/elusiv0/wb_tech_l0/internal/dto/order"
)

type OrderRepo interface {
	GetAll(ctx context.Context) ([]order.Order, error)
	Insert(ctx context.Context, order order.Order) (order.Order, error)
}
