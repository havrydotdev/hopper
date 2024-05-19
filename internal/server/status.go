package server

import (
	"io"
	"net"

	cbound "github.com/gavrylenkoIvan/hopper/public/clientbound"
	sbound "github.com/gavrylenkoIvan/hopper/public/serverbound"
	"github.com/gavrylenkoIvan/hopper/public/types"
)

func (h *Hopper) status(conn net.Conn) error {
	for i := 0; i < 2; i++ {
		_, packetID, err := ReadPacketInfo(conn)
		if err != nil {
			return err
		}

		switch int(packetID) {
		case sbound.PingPacketID:
			payload := make([]byte, types.LongBytes)
			_, err = io.ReadFull(conn, payload)
			if err != nil {
				break
			}

			var body []byte
			body, err = PrependID(sbound.PingPacketID, payload)
			if err != nil {
				break
			}

			_, err = WriteRaw(conn, body)
		case cbound.ListPacketID:
			players := cbound.Players{
				Max:    h.Config.Motd.MaxPlayers,
				Online: 0,
			}

			_, err = WritePacket(conn,
				cbound.ListPacketID,
				cbound.NewList(h.Config.Motd.Description, players, h.favicon),
			)
		}

		if err != nil {
			return err
		}
	}

	return nil
}
