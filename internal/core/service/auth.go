package service

import (
	"context"

	"github.com/yehezkiel1086/go-gin-jwt-oauth2/internal/core/domain"
	"github.com/yehezkiel1086/go-gin-jwt-oauth2/internal/core/port"
	"github.com/yehezkiel1086/go-gin-jwt-oauth2/internal/core/util"
)

type AuthService struct {
	port port.UserRepository
}

func InitAuthService(port port.UserRepository) *AuthService {
	return &AuthService{
		port: port,
	}
}

func (svc *AuthService) Login(ctx context.Context, user *domain.User) (*domain.User, error) {
	// save input password
	inputPwd := user.Password

	// compare username and password
	user, err := svc.port.GetUserByUsername(ctx, user.Username)
	if err != nil {
		return &domain.User{}, err
	}

	if err := util.ComparePassword(user.Password, inputPwd); err != nil {
		return &domain.User{}, err
	}
	
	return user, nil
}
