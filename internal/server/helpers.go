package server

import (
	"bytes"

	"github.com/gavrylenkoIvan/hopper/public/types"
)

func PrependID(id int, p []byte) ([]byte, error) {
	res := bytes.NewBuffer(nil)
	_, err := types.VarInt(id).WriteTo(res)
	if err != nil {
		return nil, err
	}

	_, err = res.Write(p)

	return res.Bytes(), err
}
