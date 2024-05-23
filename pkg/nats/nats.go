package nats

import (
	"fmt"

	"github.com/elusiv0/wb_tech_l0/internal/config"
	"github.com/nats-io/stan.go"
)

func New(cfg *config.Nats) (stan.Conn, error) {
	st, err := stan.Connect(
		cfg.ClusterId,
		cfg.ClientId,
		stan.NatsURL(fmt.Sprintf("nats://%s:%s", cfg.Host, cfg.Port)),
		stan.ConnectWait(cfg.ConnectWait),
		stan.Pings(cfg.PingInterval, cfg.PingMaxOut),
	)
	if err != nil {
		return nil, fmt.Errorf("Nats - NewNats: %w", err)
	}

	return st, nil
}
