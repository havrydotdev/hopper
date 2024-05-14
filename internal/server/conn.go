package server

import (
	"log"
	"net"

	"havry.dev/havry/hopper/internal/protocol/packet"
)

const (
	StatusState int = 0x01
	LoginState  int = 0x2
)

func (h *Hopper) handleConn(conn net.Conn) error {
	defer conn.Close()

	// new conn always starts with handshake packet
	p := new(packet.Handshake)
	_, _, err := ReadPacket(conn, p)
	if err != nil {
		return err
	}

	log.Println(p)

	switch int(p.NextState) {
	case StatusState:
		return h.status(conn)
	}

	return nil
}
