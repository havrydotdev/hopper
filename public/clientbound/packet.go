package cbound

import (
	"io"

	"github.com/gavrylenkoIvan/hopper/public/interfaces"
)

type Packet interface {
	interfaces.Packet
	io.WriterTo
}
