package server

import (
	"io"
	"net"

	cbound "havry.dev/havry/hopper/internal/protocol/packets/clientbound"
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
			payload := make([]byte, types.LongBytes)
			_, err = io.ReadFull(conn, payload)
			if err != nil {
				break
			}

			var body []byte
			body, err = PrependID(PingPacketID, payload)
			if err != nil {
				break
			}

			_, err = WriteRaw(conn, body)
		case ListPacketID:
			players := cbound.Players{
				Max:    h.Config.Motd.MaxPlayers,
				Online: 0,
			}

			_, err = WritePacket(conn,
				ListPacketID,
				cbound.NewList(h.Config.Motd.Description, players, h.favicon),
			)
		}

		if err != nil {
			return err
		}
	}

	return nil
}
