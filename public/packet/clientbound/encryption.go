package cbound

import (
	"github.com/gavrylenkoIvan/hopper/public/types"
)

const (
	EncryptionID int = 0x01
)

func NewEncryption(pubKey, verifToken []byte) *Packet {
	return NewPacket(
		types.VarInt(EncryptionID),
		// ServerID
		types.String(""),
		// Public Key
		types.ByteArr(pubKey),
		// Verify Token
		types.ByteArr(verifToken),
	)
}
