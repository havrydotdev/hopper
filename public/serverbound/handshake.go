package sbound

import (
	"github.com/gavrylenkoIvan/hopper/public/types"
)

// https://wiki.vg/Protocol#Handshake
type Handshake struct {
	ProtocolVersion types.VarInt
	ServerAddress   types.String
	ServerPort      types.UShort
	NextState       types.VarInt
}
