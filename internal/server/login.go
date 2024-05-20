package server

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"log/slog"

	"github.com/gavrylenkoIvan/hopper/internal/hopper"
	cbound "github.com/gavrylenkoIvan/hopper/public/clientbound"
	sbound "github.com/gavrylenkoIvan/hopper/public/serverbound"
	"github.com/google/uuid"
)

// TODO: implement login sequence
func (h *Hopper) login(conn *hopper.Conn) error {
	// Read "Login Start" packet
	// https://wiki.vg/Protocol#Login_Start
	loginStart := new(sbound.LoginStart)
	_, _, err := conn.ReadPacket(loginStart)
	if err != nil {
		return err
	}

	slog.Info("Login start",
		slog.String("name", string(loginStart.Name)),
		slog.String("uuid", uuid.UUID(loginStart.PlayerUUID).String()),
	)

	encryption, err := cbound.NewEncryption(h.pubKey)
	if err != nil {
		return err
	}

	_, err = conn.WritePacket(encryption)
	if err != nil {
		return err
	}

	encryptionResp := new(sbound.EncryptionResp)
	_, _, err = conn.ReadPacket(encryptionResp)
	if err != nil {
		return err
	}

	slog.Debug("Encryption Response Accepted",
		slog.Any("resp", encryptionResp),
	)

	verifyToken, err := rsa.DecryptPKCS1v15(rand.Reader, h.privKey, encryptionResp.VerifToken)
	if err != nil {
		return err
	}

	if !bytes.Equal(verifyToken, encryption.VerifToken) {
		return errors.New("login: verify token does not match")
	}

	sharedSecret, err := rsa.DecryptPKCS1v15(rand.Reader, h.privKey, encryptionResp.SharedSecret)
	if err != nil {
		return err
	}

	conn.SetSharedSecret(sharedSecret)

	return nil
}
