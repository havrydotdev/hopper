package sbound

import (
	"io"

	"github.com/gavrylenkoIvan/hopper/public/types"
)

type Handshake struct {
	ProtocolVersion types.VarInt
	ServerAddress   types.String
	ServerPort      types.UShort
	NextState       types.VarInt
}

func (h *Handshake) ReadFrom(r io.Reader) (int64, error) {
	protocolVerN, err := h.ProtocolVersion.ReadFrom(r)
	if err != nil {
		return 0, err
	}

	serverAddrN, err := h.ServerAddress.ReadFrom(r)
	if err != nil {
		return 0, err
	}

	serverPortN, err := h.ServerPort.ReadFrom(r)
	if err != nil {
		return 0, err
	}

	nextStateN, err := h.NextState.ReadFrom(r)
	if err != nil {
		return 0, err
	}

	return protocolVerN + serverAddrN + serverPortN + nextStateN, nil
}
