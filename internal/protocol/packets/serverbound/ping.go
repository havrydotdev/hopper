package sbound

import (
	"io"

	"havry.dev/havry/hopper/internal/protocol/types"
)

type Ping struct {
	Payload types.Long
}

func (s *Ping) ReadFrom(r io.Reader) (int64, error) {
	return s.Payload.ReadFrom(r)
}
