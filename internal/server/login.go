package server

import (
	"bytes"
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

// TODO: add connection compression
func (h *Hopper) login(conn *hopper.Conn) error {
	// Read Login Start packet
	// https://wiki.vg/Protocol#Login_Start
	loginStart, err := h.readLoginStart(conn)
	if err != nil {
		return fmt.Errorf("loginStart: %s", err.Error())
	}

	// Send encryption request packet
	// https://wiki.vg/Protocol#Encryption_Request
	verifyToken, err := h.writeEncryptionReq(conn)
	if err != nil {
		return fmt.Errorf("encryption: %s", err.Error())
	}

	// Read encryption response packet
	// https://wiki.vg/Protocol#Encryption_Response
	encryptionResp, err := h.readEncryptionResp(conn)
	if err != nil {
		return fmt.Errorf("encryptionResp: %s", err.Error())
	}

	// Check if encryption response's verify
	// token is equal to generated one
	if !bytes.Equal(verifyToken, encryptionResp.VerifyToken) {
		return errors.New("verify token is invalid")
	}

	// Set connection encryption
	err = conn.SetEncryption(encryptionResp.SharedSecret)
	if err != nil {
		return fmt.Errorf("setEncryption: %s", err.Error())
	}

	// Make hasJoined request to mojang sessions server
	// https://wiki.vg/Protocol_Encryption#Server
	hasJoinedResp, err := h.hasJoined(string(loginStart.Name), encryptionResp.SharedSecret)
	if err != nil {
		return fmt.Errorf("hasJoined: %s", err.Error())
	}

	// Write Login Success packet
	// https://wiki.vg/Protocol#Login_Success
	err = h.writeLoginSuccess(conn, hasJoinedResp)
	if err != nil {
		return fmt.Errorf("loginSuccess: %s", err.Error())
	}

	// Read Login Acknowledged packet
	// https://wiki.vg/Protocol#Login_Acknowledged
	err = h.readLoginAcknowledged(conn)
	if err != nil {
		return fmt.Errorf("loginAcknowledged: %s", err.Error())
	}

	slog.Debug("Login acknowledged")

	return nil
}

func (h *Hopper) readLoginStart(conn *hopper.Conn) (*sbound.LoginStart, error) {
	// Read "Login Start" packet
	// https://wiki.vg/Protocol#Login_Start
	loginStart := new(sbound.LoginStart)
	_, _, err := conn.ReadPacket(
		&loginStart.Name,
		&loginStart.PlayerUUID,
	)
	if err != nil {
		return nil, err
	}

	slog.Info("Login start",
		slog.String("name", string(loginStart.Name)),
		slog.String("uuid", uuid.UUID(loginStart.PlayerUUID).String()),
	)

	return loginStart, nil
}

// Sends encryption packet and
// returns generated verify token
func (h *Hopper) writeEncryptionReq(conn *hopper.Conn) ([]byte, error) {
	verifyToken, err := helpers.NewVerifyToken()
	if err != nil {
		return nil, err
	}

	encryption, err := cbound.NewEncryption(h.pubKey, verifyToken)
	if err != nil {
		return nil, err
	}

	_, err = conn.WritePacket(encryption)

	return verifyToken, err
}

func (h *Hopper) readEncryptionResp(conn *hopper.Conn) (*sbound.EncryptionResp, error) {
	p := new(sbound.EncryptionResp)
	_, _, err := conn.ReadPacket(
		&p.SharedSecret,
		&p.VerifyToken,
	)
	if err != nil {
		return nil, err
	}

	// Decrypt encryption response packet's fields
	// https://wiki.vg/Protocol_Encryption
	err = p.Decrypt(h.privKey)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (h *Hopper) hasJoined(username string, sharedSecret []byte) (*mojang.HasJoinedResponse, error) {
	hasJoinedResp, err := mojang.HasJoined(username, sharedSecret, h.pubKey)
	if err != nil {
		return nil, err
	}

	return hasJoinedResp, nil
}

func (h *Hopper) writeLoginSuccess(conn *hopper.Conn, hasJoinedResp *mojang.HasJoinedResponse) error {
	p, err := cbound.NewLoginSuccess(hasJoinedResp)
	if err != nil {
		return err
	}

	_, err = conn.WritePacket(p)

	return err
}

func (h *Hopper) readLoginAcknowledged(conn *hopper.Conn) error {
	_, _, err := conn.ReadPacketInfo()

	return err
}
