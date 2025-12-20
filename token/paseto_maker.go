package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"

	"github.com/o1egl/paseto"
)

// paseto token maker
type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

// initializes paseto token maker
func NewPasetoMaker(symmetricKey string) (Maker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("size of the key must be exactly %d characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}
	return maker, nil
}

// username and expiry duration is encoded in to a token during creation
func (maker *PasetoMaker) CreateToken(username string, duration time.Duration, tokenType string) (string, *Payload, error) {
	payload, err := NewPayload(username, duration, tokenType)
	if err != nil {
		return "", nil, err
	}

	token, err := maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
	if err != nil {
		return "", nil, err
	}
	return token, payload, err
}

// verify token will take the token check its validity, and return the paylod
func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
