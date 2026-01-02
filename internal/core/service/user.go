package service

import (
	"context"

	"github.com/yehezkiel1086/go-gin-nextjs-auth/internal/core/domain"
	"github.com/yehezkiel1086/go-gin-nextjs-auth/internal/core/port"
	"github.com/yehezkiel1086/go-gin-nextjs-auth/internal/core/util"
)

type UserService struct {
	repo port.UserRepository
}

func NewUserService(repo port.UserRepository) (*UserService) {
	return &UserService{
		repo,
	}
}

func (us *UserService) RegisterUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	hashedPwd, err := util.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hashedPwd

	return us.repo.CreateUser(ctx, user)
}

func (us *UserService) GetUsers(ctx context.Context) ([]domain.User, error) {
	return us.repo.GetUsers(ctx)
}
