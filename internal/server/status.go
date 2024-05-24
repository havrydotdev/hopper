package server

import (
	"errors"

	"github.com/gavrylenkoIvan/hopper/internal/hopper"
	cbound "github.com/gavrylenkoIvan/hopper/public/clientbound"
	"github.com/gavrylenkoIvan/hopper/public/packet"
	sbound "github.com/gavrylenkoIvan/hopper/public/serverbound"
	"github.com/gavrylenkoIvan/hopper/public/types"
)

func (h *Hopper) status(conn *hopper.Conn) error {
	for i := 0; i < 2; i++ {
		_, packetID, err := conn.ReadPacketInfo()
		if err != nil {
			return err
		}

		respBody, err := h.getStatusResp(conn, int(packetID))
		if err != nil {
			return err
		}

		_, err = conn.WritePacket(respBody)
		if err != nil {
			return err
		}
	}

	return nil
}

func (h *Hopper) getStatusResp(conn *hopper.Conn, packetID int) ([]byte, error) {
	switch packetID {
	case sbound.PingPacketID:
		var payload types.Long
		_, err := payload.ReadFrom(conn)
		if err != nil {
			return nil, err
		}

		return packet.Marshal(
			types.VarInt(sbound.PingPacketID),
			payload,
		)

	case cbound.ListPacketID:
		players := cbound.Players{
			Max:    h.Config.Motd.MaxPlayers,
			Online: 0,
		}

		return cbound.NewList(h.Config.Motd.Description, players, h.favicon)
	}

	return nil, errors.New("unknown packet id")
}
