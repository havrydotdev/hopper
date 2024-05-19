package server

import (
	"crypto/x509"
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

	encrypted, err := x509.MarshalPKIXPublicKey(&h.pubKey)
	if err != nil {
		return err
	}

	encryption, err := cbound.NewEncryption(encrypted)
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

	return nil
}
