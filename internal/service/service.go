package service

import (
	"context"

	orderDto "github.com/elusiv0/wb_tech_l0/internal/dto/order"
)

type OrderService interface {
	AddOrder(context.Context, orderDto.Order) (orderDto.Order, error)
	GetAll(context.Context) ([]orderDto.Order, error)
}
