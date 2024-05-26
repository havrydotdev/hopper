package server

import (
	"errors"

	"github.com/gavrylenkoIvan/hopper/internal/hopper"
	cbound "github.com/gavrylenkoIvan/hopper/public/packet/clientbound"
	sbound "github.com/gavrylenkoIvan/hopper/public/packet/serverbound"
	"github.com/gavrylenkoIvan/hopper/public/types"
)

// Handle status packet
// https://wiki.vg/Protocol#Status
func (h *Hopper) handleStatus(conn *hopper.Conn) error {
	for i := 0; i < 2; i++ {
		_, packetID, err := conn.ReadPacketInfo()
		if err != nil {
			return err
		}

		respBody, err := h.getStatusResp(conn, int(packetID))
		if err != nil {
			return err
		}

		_, err = respBody.WriteTo(conn)
		if err != nil {
			return err
		}
	}

	return nil
}

func (h *Hopper) getStatusResp(conn *hopper.Conn, packetID int) (*cbound.Packet, error) {
	switch packetID {
	//https://wiki.vg/Server_List_Ping#Ping_Request
	case sbound.PingPacketID:
		var payload types.Long
		_, err := payload.ReadFrom(conn)
		if err != nil {
			return nil, err
		}

		return cbound.NewPacket(
			types.VarInt(sbound.PingPacketID),
			payload,
		), nil

	// https://wiki.vg/Server_List_Ping#Status_Response
	case cbound.ListPacketID:
		players := cbound.Players{
			Max:    h.Config.Motd.MaxPlayers,
			Online: 0,
		}

		return cbound.NewList(
			h.Config.Motd.Description,
			players,
			h.favicon,
		), nil
	}

	return nil, errors.New("unknown packet id")
}
