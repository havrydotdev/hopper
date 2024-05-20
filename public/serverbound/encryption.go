package sbound

import (
	"io"

	"github.com/gavrylenkoIvan/hopper/public/types"
)

type EncryptionResp struct {
	SharedSecretLen types.VarInt
	SharedSecret    types.ByteArr

	VerifTokenLen types.VarInt
	VerifToken    types.ByteArr
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

	verifyTokenLenN, err := e.VerifTokenLen.ReadFrom(r)
	if err != nil {
		return 0, err
	}

	e.VerifToken = types.ByteArr(make([]byte, e.SharedSecretLen))
	verifyTokenN, err := e.VerifToken.ReadFrom(r)

	return sharedSecretLenN + sharedSecretN + verifyTokenLenN + verifyTokenN, err
}
