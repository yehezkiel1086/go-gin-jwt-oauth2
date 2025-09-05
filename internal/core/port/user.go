package port

import (
	"context"

	"github.com/yehezkiel1086/go-gin-jwt-oauth2/internal/core/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUserByUsername(ctx context.Context, username string) (*domain.User, error)
	GetAllUsers(ctx context.Context) ([]domain.User, error)
}

type UserService interface {
	Register(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUserByUsername(ctx context.Context, user *domain.User) (*domain.User, error)
	GetAllUsers(ctx context.Context) ([]domain.User, error)
}
