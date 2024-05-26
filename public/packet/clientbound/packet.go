package cbound

import (
	"io"

	"github.com/gavrylenkoIvan/hopper/public/packet"
	"github.com/gavrylenkoIvan/hopper/public/types"
)

type Packet struct {
	fields []io.WriterTo
}

func NewPacket(fields ...io.WriterTo) *Packet {
	return &Packet{fields}
}

// Write packet's contents prepended with it's length into writer
func (p *Packet) WriteTo(w io.Writer) (int64, error) {
	// Marshal packet's contents
	marshaled, err := packet.Marshal(p.fields...)
	if err != nil {
		return 0, err
	}

	// Write packet's length into writer
	_, err = types.VarInt(len(marshaled)).WriteTo(w)
	if err != nil {
		return 0, err
	}

	n, err := w.Write(marshaled)

	return int64(n), err
}
