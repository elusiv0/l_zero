package order

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/elusiv0/wb_tech_l0/internal/cache"
	"github.com/elusiv0/wb_tech_l0/internal/dto/order"
	orderDto "github.com/elusiv0/wb_tech_l0/internal/dto/order"
	"github.com/elusiv0/wb_tech_l0/internal/service"
	"github.com/go-playground/validator"
	"github.com/nats-io/stan.go"
)

type OrderConsumer struct {
	stan         stan.Conn
	orderService service.OrderService
	logger       *slog.Logger
	subscription stan.Subscription
	validator    *validator.Validate
	cache        *cache.Cache
}

const (
	orders = "orders"
)

func New(
	orderService service.OrderService,
	logger *slog.Logger,
	stan stan.Conn,
	cache *cache.Cache,
) (*OrderConsumer, error) {
	orderConsumer := &OrderConsumer{
		stan:         stan,
		orderService: orderService,
		logger:       logger,
		cache:        cache,
		validator:    validator.New(),
	}
	return orderConsumer, nil
}

func (orderConsumer *OrderConsumer) Subscribe() error {
	var err error
	orderConsumer.subscription, err = orderConsumer.stan.Subscribe(
		orders,
		orderConsumer.addOrder,
		stan.DurableName("my-durable"),
		stan.DeliverAllAvailable(),
	)
	if err != nil {
		return fmt.Errorf("subscribe: %w", err)
	}

	return nil
}

func (orderConsumer *OrderConsumer) addOrder(msg *stan.Msg) {
	var order orderDto.Order
	var err error
	ctx := context.Background()
	if err = orderConsumer.unmarshallData(msg.Data, &order); err != nil {
		orderConsumer.logger.Error("OrderConsumer - addOrder - unmarshallData", slog.Any("error", err.Error()))
	}

	orderResp, err := orderConsumer.orderService.AddOrder(ctx, order)
	if err != nil {
		orderConsumer.logger.Error("OrderConsumer - addOrder " + err.Error())
		return
	}

	if err = orderConsumer.cache.AddOrder(ctx, orderResp); err != nil {
		orderConsumer.logger.Error("OrderConsumer - addOrder " + err.Error())
		return
	}
}

func (orderConsumer *OrderConsumer) unmarshallData(data []byte, order *order.Order) error {
	if err := json.Unmarshal(data, order); err != nil {
		return fmt.Errorf("Error while unmarshalling json: %w", err)
	}
	if err := orderConsumer.validator.Struct(order); err != nil {
		return fmt.Errorf("Error while validating json: %w", err)
	}

	return nil
}
