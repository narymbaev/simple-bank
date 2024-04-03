package token

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/o1egl/paseto"
	"time"
)

var (
	ErrInvalidToken       = errors.New("token is invalid")
	ErrTokenInvalidClaims = errors.New("token has invalid claims")
	ErrTokenExpired       = errors.New("token is expired")
	ErrPasetoTokenExpired = errors.New("token has expired: token validation error")
)

const (
	PasetoGenerator = "PASETO"
	JWTGenerator    = "JWT"
)

type Payload interface{}

type JWTPayload struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	jwt.RegisteredClaims
}

type PASETOPayload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	JSONToken paseto.JSONToken
}

func NewPayload(generator string, username string, duration time.Duration) (Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	if generator == JWTGenerator {
		payload := &JWTPayload{
			ID:       tokenID,
			Username: username,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: &jwt.NumericDate{
					Time: time.Now().Add(duration),
				},
				IssuedAt: &jwt.NumericDate{Time: time.Now()},
			},
		}

		return payload, nil
	} else if generator == PasetoGenerator {
		payload := &PASETOPayload{
			ID:       tokenID,
			Username: username,
			JSONToken: paseto.JSONToken{
				Audience:   "test",
				Issuer:     "test",
				Jti:        "123",
				Subject:    "123",
				Expiration: time.Now().Add(duration),
				IssuedAt:   time.Now(),
				NotBefore:  time.Now(),
			},
		}
		return payload, nil
	} else {
		return nil, fmt.Errorf("Undefined payload")
	}
}

func (p *JWTPayload) GetExpirationTime() (*jwt.NumericDate, error) {
	if time.Now().After(p.ExpiresAt.Time) {
		return nil, ErrTokenExpired
	}

	return p.ExpiresAt, nil
}

func (p *PASETOPayload) Validate() error {
	if time.Now().After(p.JSONToken.Expiration) {
		return ErrTokenExpired
	}
	return nil
}
