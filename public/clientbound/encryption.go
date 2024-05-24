package cbound

import (
	"github.com/gavrylenkoIvan/hopper/public/packet"
	"github.com/gavrylenkoIvan/hopper/public/types"
)

const (
	EncryptionID int = 0x01
)

func NewEncryption(pubKey, verifToken []byte) ([]byte, error) {
	return packet.Marshal(
		types.VarInt(EncryptionID),
		// ServerID
		types.String(""),
		// Public Key
		types.ByteArr(pubKey),
		// Verify Token
		types.ByteArr(verifToken),
	)
}
