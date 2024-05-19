package protocol

import (
	"crypto/rand"
	"crypto/rsa"
)

const PubKeyBits = 1024

func GenPubKey() (rsa.PublicKey, error) {
	key, err := rsa.GenerateKey(rand.Reader, PubKeyBits)

	return key.PublicKey, err
}
