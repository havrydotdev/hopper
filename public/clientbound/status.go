package cbound

import (
	"encoding/json"
	"io"

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
) *List {
	list := new(List)
	list.Players = players
	list.Favicon = favicon
	list.Description.Text = desc
	list.Version.Name = version
	list.Version.Protocol = protocol

	return list
}

func (l *List) ID() int {
	return ListPacketID
}

func (l *List) WriteTo(w io.Writer) (int64, error) {
	buf, err := json.Marshal(l)
	if err != nil {
		return 0, err
	}

	return types.String(buf).WriteTo(w)
}
