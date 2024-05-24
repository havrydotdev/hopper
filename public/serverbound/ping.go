package sbound

import (
	"github.com/gavrylenkoIvan/hopper/public/types"
)

const (
	PingPacketID int = 0x01
)

type Ping struct {
	Payload types.Long
}
