package port

import (
	"context"

	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/core/domain"
)

type AuthService interface {
	Login(ctx context.Context, input *domain.User) (*domain.User, error)
}
