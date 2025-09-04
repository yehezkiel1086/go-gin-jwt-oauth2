package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

	// compare username and password, generate jwt and set to cookie
	_, err := ah.port.Login(c, user)
	if err != nil {
		if err := c.BindJSON(&input); err != nil {
			util.ResponseHandler(c, http.StatusBadRequest, true, "Invalid username or password")
			return
		}
	}

	// loggedin response
	util.ResponseHandler(c, http.StatusOK, false, "Login success.")
}
