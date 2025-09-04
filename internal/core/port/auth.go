package port

import (
	"context"

	"github.com/yehezkiel1086/go-gin-jwt-oauth2/internal/core/domain"
)

type AuthService interface {
	Login(ctx context.Context, user *domain.User) (*domain.User, error)
}
