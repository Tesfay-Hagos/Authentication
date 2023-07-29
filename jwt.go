package authentication

import (
	"errors"

	"github.com/golang-jwt/jwt"
)

func (maker *jWTMaker) CreateJwtToken() (string, error) {
	payload := jWTMaker{
		secretKey:  maker.secretKey,
		ID:         maker.ID,
		UserID:     maker.UserID,
		IssuedAt:   maker.IssuedAt,
		AdminGrant: maker.AdminGrant,
		ExpiredAt:  maker.ExpiredAt,
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte(maker.secretKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (maker *jWTMaker) VerifyJwtToken(token string) (*jWTMaker, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &jWTMaker{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*jWTMaker)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
