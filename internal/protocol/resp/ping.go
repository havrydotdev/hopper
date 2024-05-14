package resp

import (
	"bytes"

	"havry.dev/havry/hopper/internal/protocol/types"
)

type Ping struct {
	Payload types.Long
}

func NewPing(payload int64) *Ping {
	return &Ping{types.Long(payload)}
}

func (p *Ping) MarshalBinary() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	_, err := p.Payload.WriteTo(buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
