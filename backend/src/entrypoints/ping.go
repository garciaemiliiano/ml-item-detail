package entrypoints

import (
	"item-detail-api/src/entrypoints/rest/ping"
)

type PingContainer struct {
	Ping ping.PingHandler
}

func NewPingContainer(Ping ping.PingHandler) PingContainer {
	return PingContainer{Ping: Ping}
}
