package sbound

import (
	"io"

	"github.com/gavrylenkoIvan/hopper/public/types"
)

const (
	PingPacketID int = 0x01
)

type Ping struct {
	Payload types.Long
}

func (s *Ping) ReadFrom(r io.Reader) (int64, error) {
	return s.Payload.ReadFrom(r)
}
