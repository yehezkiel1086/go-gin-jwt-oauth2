package service

import (
	"context"

	"github.com/yehezkiel1086/go-gin-jwt-oauth2/internal/core/domain"
	"github.com/yehezkiel1086/go-gin-jwt-oauth2/internal/core/port"
	"github.com/yehezkiel1086/go-gin-jwt-oauth2/internal/core/util"
)

type UserService struct {
	repo port.UserRepository
}

func InitUserService(repo port.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (us *UserService) Register(ctx context.Context, user *domain.User) (*domain.User, error) {
	hashedPwd, err := util.HashPassword(user.Password)
	if err != nil {
		return &domain.User{}, err
	}

	user.Password = hashedPwd

	user, err = us.repo.CreateUser(ctx, user)
	if err != nil {
		return &domain.User{}, err
	}

	return user, nil
}

func (us *UserService) GetUserByUsername(ctx context.Context, user *domain.User) (*domain.User, error) {
	user, err := us.repo.GetUserByUsername(ctx, user.Username)
	if err != nil {
		return &domain.User{}, err
	}

	return user, nil
}

func (us *UserService) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	user, err := us.repo.GetAllUsers(ctx)
	if err != nil {
		return []domain.User{}, err
	}

	return user, nil
}
