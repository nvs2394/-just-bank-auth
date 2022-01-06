package service

import (
	"github.com/golang-jwt/jwt"
	"github.com/nvs2394/just-bank-auth/domain"
	"github.com/nvs2394/just-bank-auth/dto"
	"github.com/nvs2394/just-bank-lib/errs"
	"github.com/nvs2394/just-bank-lib/logger"
)

type AuthService interface {
	Login(dto.LoginRequest) (*dto.LoginResponse, *errs.AppError)
	Verify(urlParams map[string]string) *errs.AppError
}

type DefaultAuthService struct {
	repo            domain.AuthRepositoryDb
	rolePermissions domain.RolePermissions
}

func (auth DefaultAuthService) Login(loginRequest dto.LoginRequest) (*dto.LoginResponse, *errs.AppError) {
	login, err := auth.repo.FindBy(loginRequest.Username, loginRequest.Password)
	if err != nil {
		return nil, err
	}

	claims := login.ClaimsForAccessToken()
	authToken := domain.NewAuthToken(claims)
	accessToken, err := authToken.NewAccessToken()
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{AccessToken: accessToken}, nil
}

func (auth DefaultAuthService) Verify(urlParams map[string]string) *errs.AppError {
	jwtToken, err := getJwtTokenFromString(urlParams["token"])
	if err != nil {
		return errs.NewUnauthorizedError(err.Error())
	}

	if jwtToken.Valid {
		claims := jwtToken.Claims.(*domain.AccessTokenClaims)

		if claims.IsUserRole() {
			//TODO: do something
			return nil
		}

		isAuthorized := auth.rolePermissions.IsAuthorizedFor(claims.Role, urlParams["routeName"])
		if !isAuthorized {
			return errs.NewUnauthorizedError("Role is not authorized")
		}
		return nil
	}

	return errs.NewUnauthorizedError("Invalid token")

}

func getJwtTokenFromString(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &domain.AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(domain.HMAC_SAMPLE_SECRET), nil
	})

	if err != nil {
		logger.Error("Error while parsing token: " + err.Error())
		return nil, err
	}

	return token, nil

}

func NewAuthService(repo domain.AuthRepositoryDb, permission domain.RolePermissions) DefaultAuthService {
	return DefaultAuthService{repo, permission}
}
