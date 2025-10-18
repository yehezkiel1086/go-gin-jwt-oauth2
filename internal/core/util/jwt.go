package util

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/adapter/config"
	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/core/domain"
)

func GenerateJWT(conf *config.JWT, user *domain.User) (string, int, error) {
	// convert token duration to int
	duration, err := strconv.Atoi(conf.Duration)
	if err != nil {
		return "", -1, errors.New("invalid token duration from .env")
	}

	// Create claims with multiple fields populated
	claims := domain.JWT{
		Email: user.Email,
		Role:  user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(duration) * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(conf.Secret))
	if err != nil {
		return "", -1, err
	}

	return ss, duration, nil
}
