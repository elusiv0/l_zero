package order

import (
	"context"
	"fmt"
	"log/slog"

	orderDto "github.com/elusiv0/wb_tech_l0/internal/dto/order"
	"github.com/elusiv0/wb_tech_l0/internal/repo"
	"github.com/elusiv0/wb_tech_l0/internal/service"
)

type OrderService struct {
	orderRepository repo.OrderRepo
	logger          *slog.Logger
}

var _ service.OrderService = &OrderService{}

func New(
	orderRepository repo.OrderRepo,
	logger *slog.Logger,
) *OrderService {
	service := &OrderService{
		orderRepository: orderRepository,
		logger:          logger,
	}

	return service
}

func (orderService *OrderService) AddOrder(ctx context.Context, order orderDto.Order) (orderDto.Order, error) {
	orderResp, err := orderService.orderRepository.Insert(ctx, order)
	if err != nil {
		return orderResp, fmt.Errorf("OrderService - AddOrder: %w", err)
	}

	return orderResp, nil
}

func (orderService *OrderService) GetAll(ctx context.Context) ([]orderDto.Order, error) {
	ordersResp, err := orderService.orderRepository.GetAll(ctx)
	if err != nil {
		return ordersResp, fmt.Errorf("OrderRepositroy - GetAll: %w", err)
	}

	return ordersResp, nil
}
