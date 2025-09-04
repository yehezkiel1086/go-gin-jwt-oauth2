package handler

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/yehezkiel1086/go-gin-jwt-oauth2/internal/core/domain"
	"github.com/yehezkiel1086/go-gin-jwt-oauth2/internal/core/port"
	"github.com/yehezkiel1086/go-gin-jwt-oauth2/internal/core/util"
)

type AuthHandler struct {
	port port.AuthService
}

func InitAuthHandler(port port.AuthService) *AuthHandler {
	return &AuthHandler{
		port: port,
	}
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Claims struct {
	Username string `json:"username"`
	Role domain.Role `json:"role"`
	jwt.RegisteredClaims
}

func (ah *AuthHandler) Login(c *gin.Context) {
	// get request
	var input LoginInput

	if err := c.BindJSON(&input); err != nil {
		util.ResponseHandler(c, http.StatusBadRequest, true, "Username and password are required")
		return
	}

	// create user object
	user := &domain.User{
		Username: input.Username,
		Password: input.Password,
	}

	// compare username and password
	user, err := ah.port.Login(c, user)
	if err != nil {
		util.ResponseHandler(c, http.StatusUnauthorized, true, "Invalid username or password")
		return
	}

	// get token secret
	secret := os.Getenv("JWT_SECRET")
	duration, err := strconv.Atoi(os.Getenv("TOKEN_DURATION"))
	if err != nil {
		util.ResponseHandler(c, http.StatusInternalServerError, true, err.Error())
		return
	}

	expDate := time.Now().Add(time.Duration(duration) * time.Minute)

	// generate jwt
	claims := Claims{
		Username: user.Username,
		Role: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			// Also fixed dates can be used for the NumericDate
			ExpiresAt: jwt.NewNumericDate(expDate),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(secret))
	if err != nil {
		util.ResponseHandler(c, http.StatusInternalServerError, true, "Failed to sign token.")
		return
	}

	// set jwt to cookie
	c.SetCookie("jwt_token", ss, duration * 60, "/", "", false, true)

	// loggedin response
	util.ResponseHandler(c, http.StatusOK, false, "Login success.")
}
