package cbound

import (
	"encoding/json"

	"github.com/gavrylenkoIvan/hopper/public/packet"
	"github.com/gavrylenkoIvan/hopper/public/types"
)

const (
	ListPacketID int = 0x00

	version  = "1.20.4"
	protocol = 765
)

type Players struct {
	Max    uint `json:"max"`
	Online int  `json:"online"`

	Sample []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"sample"`
}

// https://wiki.vg/Server_List_Ping#Status_Response
type List struct {
	Players Players `json:"players"`
	Favicon *string `json:"favicon"`

	Description struct {
		Text string `json:"text"`
	} `json:"description"`

	Version struct {
		Name     string `json:"name"`
		Protocol int    `json:"protocol"`
	} `json:"version"`
}

func NewList(
	desc string,
	players Players,
	favicon *string,
) ([]byte, error) {
	list := new(List)
	list.Players = players
	list.Favicon = favicon
	list.Description.Text = desc
	list.Version.Name = version
	list.Version.Protocol = protocol

	buf, err := json.Marshal(list)
	if err != nil {
		return nil, err
	}

	return packet.Marshal(
		types.VarInt(ListPacketID),
		types.String(buf),
	)
}
