package sbound

import (
	"crypto/rand"
	"crypto/rsa"

	"github.com/gavrylenkoIvan/hopper/public/types"
)

// https://wiki.vg/Protocol#Encryption_Response
type EncryptionResp struct {
	SharedSecret types.ByteArr
	VerifyToken  types.ByteArr
}

func (e *EncryptionResp) Decrypt(privateKey *rsa.PrivateKey) error {
	var err error
	e.SharedSecret, err = rsa.DecryptPKCS1v15(rand.Reader, privateKey, e.SharedSecret)
	if err != nil {
		return err
	}

	e.VerifyToken, err = rsa.DecryptPKCS1v15(rand.Reader, privateKey, e.VerifyToken)

	return err
}
