package server

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"fmt"
	"log/slog"

	"github.com/gavrylenkoIvan/hopper/internal/hopper"
	cbound "github.com/gavrylenkoIvan/hopper/public/clientbound"
	"github.com/gavrylenkoIvan/hopper/public/helpers"
	"github.com/gavrylenkoIvan/hopper/public/mojang"
	sbound "github.com/gavrylenkoIvan/hopper/public/serverbound"
	"github.com/google/uuid"
)

// TODO: implement login sequence
func (h *Hopper) login(conn *hopper.Conn) error {
	// Read "Login Start" packet
	// https://wiki.vg/Protocol#Login_Start
	var loginStart sbound.LoginStart
	_, _, err := conn.ReadPacket(
		&loginStart.Name,
		&loginStart.PlayerUUID,
	)
	if err != nil {
		return fmt.Errorf("loginStart: %s", err.Error())
	}

	slog.Info("Login start",
		slog.String("name", string(loginStart.Name)),
		slog.String("uuid", uuid.UUID(loginStart.PlayerUUID).String()),
	)

	verifToken, err := helpers.NewVerifToken()
	if err != nil {
		return err
	}

	encryption, err := cbound.NewEncryption(h.pubKey, verifToken)
	if err != nil {
		return fmt.Errorf("encryption: %s", err.Error())
	}

	_, err = conn.WritePacket(encryption)
	if err != nil {
		return fmt.Errorf("write encryption: %s", err.Error())
	}

	var encryptionResp sbound.EncryptionResp
	_, _, err = conn.ReadPacket(
		&encryptionResp.SharedSecret,
		&encryptionResp.VerifToken,
	)
	if err != nil {
		return fmt.Errorf("encryptionResp: %s", err.Error())
	}

	verifyToken, err := rsa.DecryptPKCS1v15(rand.Reader, h.privKey, encryptionResp.VerifToken)
	if err != nil {
		return fmt.Errorf("verifyToken: %s", err.Error())
	}

	if !bytes.Equal(verifyToken, verifToken) {
		return errors.New("login: verify token does not match")
	}

	sharedSecret, err := rsa.DecryptPKCS1v15(rand.Reader, h.privKey, encryptionResp.SharedSecret)
	if err != nil {
		return fmt.Errorf("sharedSecret: %s", err.Error())
	}

	hasJoinedResp, err := mojang.HasJoined(string(loginStart.Name), sharedSecret, h.pubKey)
	if err != nil {
		return fmt.Errorf("hasJoined: %s", err.Error())
	}

	err = conn.SetEncryption(sharedSecret)
	if err != nil {
		return err
	}

	ls, err := cbound.NewLoginSuccess(hasJoinedResp)
	if err != nil {
		return err
	}

	_, err = conn.WritePacket(ls)

	return err
}
