package server

import (
	"errors"

	"github.com/gavrylenkoIvan/hopper/internal/hopper"
	sbound "github.com/gavrylenkoIvan/hopper/public/packet/serverbound"
)

const (
	StatusState int = 0x01
	LoginState  int = 0x2
)

func (h *Hopper) handshake(conn *hopper.Conn) error {
	defer conn.Close()

	// new conn always starts with handshake packet
	var p sbound.Handshake
	_, _, err := conn.ReadPacket(
		&p.ProtocolVersion,
		&p.ServerAddress,
		&p.ServerPort,
		&p.NextState,
	)
	if err != nil {
		return err
	}

	switch int(p.NextState) {
	case StatusState:
		return h.handleStatus(conn)
	case LoginState:
		return h.handleLogin(conn)
	}

	return errors.New("unknown packet")
}
