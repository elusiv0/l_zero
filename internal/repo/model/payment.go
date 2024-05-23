package model

type Payment struct {
	PaymentId          int
	PaymentTransaction string
	RequiestId         string
	Currency           string
	PaymentProvider    string
	Amount             int
	PaymentDt          int
	Bank               string
	DeliveryCost       int
	GoodsTotal         int
	CustomFee          int
}
