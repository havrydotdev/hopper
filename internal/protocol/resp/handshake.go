package resp

import (
	"bytes"
	"encoding/json"

	"havry.dev/havry/hopper/internal/protocol/types"
)

// TODO refactor this shit

type Description struct {
	Text string `json:"text"`
}

type version struct {
	Name     string `json:"name"`
	Protocol int    `json:"protocol"`
}

type Players struct {
	Max    int      `json:"max"`
	Online int      `json:"online"`
	Sample []sample `json:"sample"`
}

type sample struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type HandshakeData struct {
	Version     version     `json:"version"`
	Players     Players     `json:"players"`
	Description Description `json:"description"`
	Favicon     *string     `json:"favicon"`
}

type Handshake struct {
	data types.String
}

func NewHandshake(
	players Players,
	description Description,
	favicon *string,
) (*Handshake, error) {
	handshake, err := json.Marshal(HandshakeData{
		Version: version{
			Name:     "1.20.4",
			Protocol: 765,
		},
		Players:     players,
		Description: description,
		Favicon:     favicon,
	})
	if err != nil {
		return nil, err
	}

	return &Handshake{types.String(handshake)}, nil
}

func (h *Handshake) MarshalBinary() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	_, err := h.data.WriteTo(buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
