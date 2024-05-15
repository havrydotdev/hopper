package packet

import (
	"encoding/json"
	"io"

	"havry.dev/havry/hopper/internal/protocol/types"
)

const (
	version  = "1.20.4"
	protocol = 765
)

type Ping struct {
	Payload types.Long
}

func (s *Ping) ReadFrom(r io.Reader) (int64, error) {
	return s.Payload.ReadFrom(r)
}

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

func (l *List) WriteTo(w io.Writer) (int64, error) {
	buf, err := json.Marshal(l)
	if err != nil {
		return 0, err
	}

	return types.String(buf).WriteTo(w)
}
