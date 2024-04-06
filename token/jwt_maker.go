package token

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

const minSecretKeySize = 32

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}
	return &JWTMaker{secretKey}, nil
}

func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	newPayload, err := NewPayload(JWTGenerator, username, duration)
	if err != nil {
		return "", err
	}
	payload, ok := newPayload.(*JWTPayload)
	if ok {
		jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
		return jwtToken.SignedString([]byte(maker.secretKey))
	}
	return "", fmt.Errorf("invalid Payload")
}

func (maker *JWTMaker) VerifyToken(token string) (Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}

		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &JWTPayload{}, keyFunc)

	// if token is not Valid, it will return err
	if !jwtToken.Valid {
		if strings.Contains(err.Error(), ErrTokenExpired.Error()) {
			return nil, ErrTokenExpired
		}

		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*JWTPayload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
