package sbound

import (
	"github.com/gavrylenkoIvan/hopper/public/types"
)

// https://wiki.vg/Protocol#Login_Start
type LoginStart struct {
	Name       types.String
	PlayerUUID types.UUID
}
