package cbound

import (
	"crypto/rand"
	"io"
	"log/slog"

	"github.com/gavrylenkoIvan/hopper/public/types"
)

const (
	verifTokenLen = 4

	EncryptionID int = 0x01
)

type Encryption struct {
	ServerID types.String

	PubKeyLen types.VarInt
	PubKey    types.ByteArr

	VerifTokenLen types.VarInt
	VerifToken    types.ByteArr
}

func NewEncryption(pubKey []byte) (*Encryption, error) {
	verifToken := make([]byte, verifTokenLen)
	_, err := rand.Read(verifToken)
	if err != nil {
		return nil, err
	}

	return &Encryption{
		ServerID: types.String(""),

		PubKey:    types.ByteArr(pubKey),
		PubKeyLen: types.VarInt(len(pubKey)),

		VerifToken:    types.ByteArr(verifToken),
		VerifTokenLen: verifTokenLen,
	}, nil
}

func (e *Encryption) ID() int {
	return EncryptionID
}

func (e *Encryption) WriteTo(w io.Writer) (int64, error) {
	slog.Debug("", slog.Any("encryption", e))
	serverIDN, err := e.ServerID.WriteTo(w)
	if err != nil {
		return 0, err
	}

	pubKeyLenN, err := e.PubKeyLen.WriteTo(w)
	if err != nil {
		return 0, err
	}

	pubKeyN, err := e.PubKey.WriteTo(w)
	if err != nil {
		return 0, err
	}

	verifTokenLenN, err := e.VerifTokenLen.WriteTo(w)
	if err != nil {
		return 0, err
	}

	verifTokenN, err := e.VerifToken.WriteTo(w)
	if err != nil {
		return 0, err
	}

	return serverIDN + pubKeyLenN + pubKeyN + verifTokenLenN + verifTokenN, nil
}
