package server

import (
	"net"

	sbound "havry.dev/havry/hopper/internal/protocol/packets/serverbound"
)

const (
	StatusState int = 0x01
	LoginState  int = 0x2
)

func (h *Hopper) handshake(conn net.Conn) error {
	defer conn.Close()

	// new conn always starts with handshake packet
	p := new(sbound.Handshake)
	_, _, err := ReadPacket(conn, p)
	if err != nil {
		return err
	}

	switch int(p.NextState) {
	case StatusState:
		return h.status(conn)
	case LoginState:
		return h.login(conn)
	}

	return nil
}
