package sbound

import (
	"github.com/gavrylenkoIvan/hopper/public/types"
)

type EncryptionResp struct {
	SharedSecret types.ByteArr
	VerifToken   types.ByteArr
}
