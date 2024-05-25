package sbound

import (
	"github.com/gavrylenkoIvan/hopper/public/types"
)

const (
	PingPacketID int = 0x01
)

// https://wiki.vg/Protocol#Status
type Ping struct {
	Payload types.Long
}
