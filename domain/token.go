package domain

import (
	"github.com/golang-jwt/jwt"
	"github.com/nvs2394/just-bank-auth/errs"
	"github.com/nvs2394/just-bank-auth/logger"
)

type AuthToken struct {
	token *jwt.Token
}

func NewAuthToken(claims AccessTokenClaims) AuthToken {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return AuthToken{
		token: token,
	}

}

func (auth AuthToken) NewAccessToken() (string, *errs.AppError) {
	secretBytes := []byte(HMAC_SAMPLE_SECRET)
	signedString, err := auth.token.SignedString(secretBytes)
	if err != nil {
		logger.Error("Failed while siging token " + err.Error())
		return "", errs.NewUnexpectedError("Can not generate token")
	}
	return signedString, nil
}
