package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/elusiv0/wb_tech_l0/internal/dto/order"
	"github.com/nats-io/stan.go"
)

type Orders struct {
	Orders []order.Order `json:"orders"`
}

func main() {
	var orders Orders
	dat, err := os.ReadFile("./data/orders")
	if err != nil {
		log.Fatal("error while reading file" + err.Error())
	}
	sc, _ := stan.Connect("test-cluster", "test-cli")
	json.Unmarshal(dat, &orders)
	for _, val := range orders.Orders {
		data, _ := json.Marshal(val)
		sc.Publish("orders", data)
	}
}
