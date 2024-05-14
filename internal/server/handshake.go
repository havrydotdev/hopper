package server

import (
	"net"

	"havry.dev/havry/hopper/internal/protocol/resp"
	"havry.dev/havry/hopper/internal/protocol/types"
)

const (
	ListPacketID int = 0x00
	PingPacketID int = 0x01
)

func (h *Hopper) status(conn net.Conn) error {
	for i := 0; i < 2; i++ {
		_, packetID, err := ReadPacketInfo(conn)
		if err != nil {
			return err
		}

		switch int(packetID) {
		case PingPacketID:
			var ping types.Long
			_, err = ping.ReadFrom(conn)
			if err != nil {
				return err
			}

			_, err = WriteResp(conn, PingPacketID, resp.NewPing(int64(ping)))
		case ListPacketID:
			var handshake *resp.Handshake
			handshake, err = resp.NewHandshake(resp.Players{
				Max:    20,
				Online: 0,
			}, resp.Description{
				Text: "PEREMOGA BUDEEEE URAAAA",
			}, nil)
			if err != nil {
				return err
			}

			_, err = WriteResp(conn, ListPacketID, handshake)
		}

		if err != nil {
			return err
		}
	}

	return nil
}
