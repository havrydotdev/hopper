package server

import (
	"io"

	"github.com/gavrylenkoIvan/hopper/internal/hopper"
	cbound "github.com/gavrylenkoIvan/hopper/public/clientbound"
	sbound "github.com/gavrylenkoIvan/hopper/public/serverbound"
	"github.com/gavrylenkoIvan/hopper/public/types"
)

func (h *Hopper) status(conn *hopper.Conn) error {
	for i := 0; i < 2; i++ {
		_, packetID, err := conn.ReadPacketInfo()
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

			_, err = conn.WriteRaw(body)
		case cbound.ListPacketID:
			players := cbound.Players{
				Max:    h.Config.Motd.MaxPlayers,
				Online: 0,
			}

			_, err = conn.WritePacket(
				cbound.NewList(h.Config.Motd.Description, players, h.favicon),
			)
		}

		if err != nil {
			return err
		}
	}

	return nil
}
