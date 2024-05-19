package sbound

import (
	"io"

	"havry.dev/havry/hopper/internal/protocol/types"
)

type EncryptionResp struct {
	SharedSecretLen types.VarInt
	SharedSecret    types.ByteArr

	VerifyTokenLen types.VarInt
	VerifyToken    types.ByteArr
}

func (e *EncryptionResp) ReadFrom(r io.Reader) (int64, error) {
	sharedSecretLenN, err := e.SharedSecretLen.ReadFrom(r)
	if err != nil {
		return 0, err
	}

	e.SharedSecret = types.ByteArr(make([]byte, e.SharedSecretLen))
	sharedSecretN, err := e.SharedSecret.ReadFrom(r)
	if err != nil {
		return 0, err
	}

	verifyTokenLenN, err := e.VerifyTokenLen.ReadFrom(r)
	if err != nil {
		return 0, err
	}

	e.VerifyToken = types.ByteArr(make([]byte, e.SharedSecretLen))
	verifyTokenN, err := e.VerifyToken.ReadFrom(r)

	return sharedSecretLenN + sharedSecretN + verifyTokenLenN + verifyTokenN, err
}
