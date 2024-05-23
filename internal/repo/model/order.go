package model

import "time"

type Order struct {
	OrderUuid         string
	TrackNumber       string
	OrderEntry        string
	Delivery          Delivery
	Payment           Payment
	Items             []Item
	Locale            string
	InternalSignature string
	CustomerId        string
	DeliveryService   string
	ShardKey          int
	SmId              int
	DateCreated       time.Time
	OofShard          int
}
