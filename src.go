package authentication

import (
	"errors"
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/google/uuid"
	"github.com/o1egl/paseto"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

const minSecretKeySize = 32

type JwtMaker interface {
	CreateJwtToken() (string, error)
	VerifyJwtToken(token string) (*jWTMaker, error)
}

type jWTMaker struct {
	secretKey  string
	ID         uuid.UUID `json:"id"`
	UserID     string    `json:"username"`
	IssuedAt   time.Time `json:"issued_at"`
	AdminGrant string    `json:"admingrant"`
	ExpiredAt  time.Time `json:"expired_at"`
}

func (payload jWTMaker) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

func NewJWTMaker(secretKey string, username string, admingrant string, duration time.Duration) (JwtMaker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	return &jWTMaker{
		secretKey:  secretKey,
		ID:         tokenID,
		UserID:     username,
		IssuedAt:   time.Now(),
		AdminGrant: admingrant,
		ExpiredAt:  time.Now().Add(duration),
	}, nil
}

type PassetoMaker interface {
	CreatepassetoToken() (string, error)
	VerifyPassetoToken(token string) (*pasetoMaker, error)
}

type pasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
	ID           uuid.UUID `json:"id"`
	UserID       string    `json:"username"`
	IssuedAt     time.Time `json:"issued_at"`
	AdminGrant   string    `json:"admingrant"`
	ExpiredAt    time.Time `json:"expired_at"`
}

func (payload *pasetoMaker) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
func NewPasetoMaker(symmetricKey string, username string, admingrant string, duration time.Duration) (PassetoMaker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	maker := &pasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
		ID:           tokenID,
		UserID:       username,
		IssuedAt:     time.Now(),
		AdminGrant:   admingrant,
		ExpiredAt:    time.Now().Add(duration),
	}

	return maker, nil
}
