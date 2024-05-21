package cbound

import (
	"io"

	"github.com/gavrylenkoIvan/hopper/public/mojang"
	"github.com/gavrylenkoIvan/hopper/public/types"
)

const LoginSuccessID int = 0x02

type LoginSuccess struct {
	UUID       types.UUID
	Username   types.String
	Properties types.Array[*Property]
}

type Property struct {
	Name  types.String
	Value types.String

	IsSigned  types.Boolean
	Signature types.String
}

func NewProperty(p mojang.Property) *Property {
	return &Property{
		Name:      types.String(p.Name),
		Value:     types.String(p.Value),
		IsSigned:  types.Boolean(p.Signature != ""),
		Signature: types.String(p.Signature),
	}
}

func (p *Property) WriteTo(w io.Writer) (n int64, err error) {
	nameN, err := p.Name.WriteTo(w)
	if err != nil {
		return 0, err
	}
	n += nameN

	valueN, err := p.Value.WriteTo(w)
	if err != nil {
		return 0, err
	}
	n += valueN

	isSignedN, err := p.IsSigned.WriteTo(w)
	if err != nil {
		return 0, err
	}
	n += isSignedN

	if p.IsSigned {
		signatureN, err := p.Signature.WriteTo(w)
		if err != nil {
			return 0, err
		}

		n += signatureN
	}

	return
}

func NewLoginSuccess(resp *mojang.HasJoinedResponse) *LoginSuccess {
	p := new(LoginSuccess)
	p.UUID = types.UUID(resp.ID)
	p.Username = types.String(resp.Name)

	props := make([]*Property, len(resp.Properties)-1)
	for _, prop := range resp.Properties {
		props = append(props, NewProperty(prop))
	}

	p.Properties = types.Array[*Property](props)

	return p
}

func (ls *LoginSuccess) ID() int {
	return LoginSuccessID
}

func (ls *LoginSuccess) WriteTo(w io.Writer) (int64, error) {
	uuidN, err := ls.UUID.WriteTo(w)
	if err != nil {
		return 0, err
	}

	usernameN, err := ls.Username.WriteTo(w)
	if err != nil {
		return 0, err
	}

	propsN, err := ls.Properties.WriteTo(w)

	return uuidN + usernameN + propsN, err
}
