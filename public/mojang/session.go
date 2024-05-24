package mojang

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

const hasJoinedURL = "https://sessionserver.mojang.com/session/minecraft/hasJoined"

type HasJoinedResponse struct {
	ID         uuid.UUID  `json:"id"`
	Name       string     `json:"name"`
	Properties []Property `json:"properties"`
}

type Property struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	Signature string `json:"signature"`
}

func HasJoined(username string, sharedSecret, publicKey []byte) (*HasJoinedResponse, error) {
	u, err := url.Parse(hasJoinedURL)
	if err != nil {
		return nil, err
	}

	loginHash := AuthDigest("", sharedSecret, publicKey)
	queryParams := buildQueryParams(username, loginHash)

	u.RawQuery = queryParams.Encode()

	return makeHasJoinedReq(u.String())
}

func makeHasJoinedReq(u string) (*HasJoinedResponse, error) {
	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var parsedResp HasJoinedResponse
	err = json.Unmarshal(body, &parsedResp)

	return &parsedResp, err
}

func buildQueryParams(username, loginHash string) url.Values {
	val := url.Values{}
	val.Set("username", username)
	val.Set("serverId", loginHash)

	return val
}
