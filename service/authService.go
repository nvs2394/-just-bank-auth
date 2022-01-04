package service

import (
	"github.com/nvs2394/just-bank-auth/domain"
	"github.com/nvs2394/just-bank-auth/dto"
	"github.com/nvs2394/just-bank-auth/errs"
)

type AuthService interface {
	Login(dto.LoginRequest) (*dto.LoginResponse, *errs.AppError)
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

func NewAuthService(repo domain.AuthRepositoryDb, permission domain.RolePermissions) DefaultAuthService {
	return DefaultAuthService{repo, permission}
}
