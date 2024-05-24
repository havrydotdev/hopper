package sbound

import (
	"github.com/gavrylenkoIvan/hopper/public/types"
)

type LoginStart struct {
	Name       types.String
	PlayerUUID types.UUID
}
