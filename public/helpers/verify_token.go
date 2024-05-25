package helpers

import "crypto/rand"

const verifTokenLen = 4

func NewVerifyToken() ([]byte, error) {
	verifToken := make([]byte, verifTokenLen)
	_, err := rand.Read(verifToken)
	if err != nil {
		return nil, err
	}

	return verifToken, nil
}
