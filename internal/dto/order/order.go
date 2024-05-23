package order

import (
	"time"

	"github.com/elusiv0/wb_tech_l0/internal/dto/delivery"
	"github.com/elusiv0/wb_tech_l0/internal/dto/item"
	"github.com/elusiv0/wb_tech_l0/internal/dto/payment"
)

type Order struct {
	OrderUid          string            `json:"order_uid" validate:"required"`
	TrackNumber       string            `json:"track_number" validate:"required"`
	Entry             string            `json:"entry" validate:"required"`
	Delivery          delivery.Delivery `json:"delivery" validate:"required"`
	Payment           payment.Payment   `json:"payment" validate:"required"`
	Items             []item.Item       `json:"items" validate:"required"`
	Locale            string            `json:"locale" validate:"required"`
	InternalSignature string            `json:"internal_signature"`
	CustomerId        string            `json:"customer_id" validate:"required"`
	DeliveryService   string            `json:"delivery_service" validate:"required"`
	ShardKey          int               `json:"shardkey" validate:"required"`
	SmId              int               `json:"sm_id" validate:"required"`
	DateCreated       time.Time         `json:"date_created" validate:"required"`
	OofShard          int               `json:"oof_shard" validate:"required"`
}
