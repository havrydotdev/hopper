package sbound

import (
	"io"

	"github.com/gavrylenkoIvan/hopper/public/types"
)

type LoginStart struct {
	Name       types.String
	PlayerUUID types.UUID
}

func (p *LoginStart) ReadFrom(r io.Reader) (int64, error) {
	nameN, err := p.Name.ReadFrom(r)
	if err != nil {
		return 0, err
	}

	uuidN, err := p.PlayerUUID.ReadFrom(r)
	if err != nil {
		return 0, err
	}

	return nameN + uuidN, err
}
