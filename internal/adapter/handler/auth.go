package handler

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/adapter/config"
	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/core/domain"
	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/core/port"
)

type AuthHandler struct {
	svc port.AuthService
	jwtConf *config.JWT
	httpConf *config.HTTP
}

func InitAuthHandler(svc port.AuthService, jwtConf *config.JWT, httpConf *config.HTTP) *AuthHandler {
	return &AuthHandler{
		svc: svc,
		jwtConf: jwtConf,
		httpConf: httpConf,
	}
}

type LoginReq struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (ah *AuthHandler) Login(c *gin.Context) {
	// bind inputs
	var input LoginReq
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "email and password are required",
		})
		return
	}

	// login
	user, err := ah.svc.Login(c, &domain.User{
		Email: input.Email,
		Password: input.Password,
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	// handle jwt
	// convert token duration to int
	duration, err := strconv.Atoi(ah.jwtConf.Duration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "invalid token duration from .env",
		})
		return
	}

	// Create claims with multiple fields populated
	claims := domain.JWT{
		Email: user.Email,
		Role: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(duration) * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(ah.jwtConf.Secret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// set cookie
	c.SetCookie("jwt_token", ss, duration * 60, "/", ah.httpConf.Host, os.Getenv("APP_ENV") == "production", true)

	// return ok response
	c.JSON(http.StatusOK, gin.H{
		"message": "user logged in successfully",
	})
}

func (ah *AuthHandler) Logout(c *gin.Context) {
	// remove jwt_token cookie
	c.SetCookie("jwt_token", "", -1, "/", ah.httpConf.Host, os.Getenv("APP_ENV") == "production", true)

	// return http response
	c.JSON(http.StatusOK, gin.H{
		"message": "logged out successfully",
	})
}
