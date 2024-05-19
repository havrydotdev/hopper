package server

import (
	"crypto/x509"
	"log/slog"
	"net"

	"github.com/google/uuid"
	cbound "github.com/gavrylenkoIvan/hopper/public/clientbound"
	sbound "github.com/gavrylenkoIvan/hopper/public/serverbound"
)

// TODO: implement login sequence
func (h *Hopper) login(conn net.Conn) error {
	// Read "Login Start" packet
	// https://wiki.vg/Protocol#Login_Start
	loginStart := new(sbound.LoginStart)
	_, _, err := ReadPacket(conn, loginStart)
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

	_, err = WritePacket(conn, encryption.ID(), encryption)
	if err != nil {
		return err
	}

	encryptionResp := new(sbound.EncryptionResp)
	_, _, err = ReadPacket(conn, encryptionResp)
	if err != nil {
		return err
	}

	slog.Debug("Encryption Response Accepted",
		slog.Any("resp", encryptionResp),
	)

	return nil
}
