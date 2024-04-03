package token

import (
	"fmt"
	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
	"time"
)

type PasetoMaker struct {
	paseto      *paseto.V2
	symetricKey []byte
}

func NewPasetoMaker(symmetricKey string) (Maker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be %d characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto:      paseto.NewV2(),
		symetricKey: []byte(symmetricKey),
	}

	return maker, nil
}

func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(PasetoGenerator, username, duration)
	if err != nil {
		return "", err
	}

	return maker.paseto.Encrypt(maker.symetricKey, payload, nil)
}

func (maker *PasetoMaker) VerifyToken(token string) (Payload, error) {
	payload := &PASETOPayload{}

	err := maker.paseto.Decrypt(token, maker.symetricKey, payload, nil)

	if err != nil {
		return nil, ErrInvalidToken
	}

	// To validate Expiration of the Token
	err = payload.Validate()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
