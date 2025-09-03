package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yehezkiel1086/go-gin-jwt-oauth2/internal/core/domain"
	"github.com/yehezkiel1086/go-gin-jwt-oauth2/internal/core/port"
)

type UserHandler struct {
	svc port.UserService
}

func InitUserHandler(svc port.UserService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func handleResponse(c *gin.Context, status int, msg string) {
	c.JSON(status, gin.H{
		"message": msg,
	})
}

func (uh *UserHandler) Register(c *gin.Context) {
	// bind request
	var input RegisterInput
	if err := c.BindJSON(&input); err != nil {
		handleResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// create user
	user := &domain.User{
		Username: input.Username,
		Password: input.Password,
		Fullname: input.Fullname,
		Email: input.Email,
	}

	// store user
	_, err := uh.svc.Register(c, user)
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// create user success
	handleResponse(c, http.StatusCreated, "Registration success.")
}
