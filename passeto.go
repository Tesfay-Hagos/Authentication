package authentication

func (maker *pasetoMaker) CreatepassetoToken() (string, error) {
	payload := pasetoMaker{
		ID:         maker.ID,
		UserID:     maker.UserID,
		IssuedAt:   maker.IssuedAt,
		AdminGrant: maker.AdminGrant,
		ExpiredAt:  maker.ExpiredAt,
	}
	return maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
}
func (maker *pasetoMaker) VerifyPassetoToken(token string) (*pasetoMaker, error) {
	payload := &pasetoMaker{}

	err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return &pasetoMaker{
		ID:         payload.ID,
		UserID:     payload.UserID,
		IssuedAt:   payload.IssuedAt,
		AdminGrant: payload.AdminGrant,
		ExpiredAt:  payload.ExpiredAt,
	}, nil
}
