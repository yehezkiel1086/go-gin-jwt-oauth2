package service

import (
	"context"

	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/core/domain"
	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/core/port"
	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/core/util"
)

type AuthService struct {
	userRepo port.UserRepository
}

func InitAuthService(userRepo port.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (as *AuthService) Login(ctx context.Context, input *domain.User) (*domain.User, error) {
	// get user by email (check email)
	user, err := as.userRepo.GetUserByEmail(ctx, input.Email)
	if err != nil {
		return &domain.User{}, err
	}

	// compare input and user password
	if err := util.ComparePassword(user.Password, input.Password); err != nil {
		return &domain.User{}, err
	}

	return user, nil
}
