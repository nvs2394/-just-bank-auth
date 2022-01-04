package domain

import (
	"database/sql"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

type Login struct {
	UserName   string         `db:"username"`
	CustomerId sql.NullString `db:"customer_id"`
	Accounts   sql.NullString `db:"account_numbers"`
	Role       string         `db:"role"`
}

func (login Login) ClaimsForAccessToken() AccessTokenClaims {
	if login.Accounts.Valid && login.CustomerId.Valid {
		return login.claimsForUser()
	}
	return login.claimsForAdmin()
}

func (login Login) claimsForUser() AccessTokenClaims {
	accounts := strings.Split(login.Accounts.String, ",")
	return AccessTokenClaims{
		CustomerId: login.CustomerId.String,
		Role:       login.Role,
		Username:   login.UserName,
		Accounts:   accounts,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ACCESS_TOKEN_DURATION).Unix(),
		},
	}
}

func (login Login) claimsForAdmin() AccessTokenClaims {
	return AccessTokenClaims{
		Role:     login.Role,
		Username: login.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ACCESS_TOKEN_DURATION).Unix(),
		},
	}
}
