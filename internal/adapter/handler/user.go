package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yehezkiel1086/go-gin-jwt-oauth2/internal/core/domain"
	"github.com/yehezkiel1086/go-gin-jwt-oauth2/internal/core/port"
	"github.com/yehezkiel1086/go-gin-jwt-oauth2/internal/core/util"
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

type GetUserByUsernameResponse struct {
	ID        uint `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Role domain.Role `json:"role"`
}

func (uh *UserHandler) GetUserByUsername(c *gin.Context) {
	// get username param
	username := c.Param("username")
	if username == "" {
		util.ResponseHandler(c, http.StatusBadRequest, true, "Username parameter is required.")
		return
	}

	// get user by username
	user, err := uh.svc.GetUserByUsername(c, &domain.User{
		Username: username,
	})
	if err != nil {
		util.ResponseHandler(c, http.StatusBadRequest, true, err.Error())
		return
	}

	// response
	c.JSON(http.StatusOK, &GetUserByUsernameResponse{
		ID: user.ID,
		Username: user.Username,
		Fullname: user.Fullname,
		Email: user.Email,
		Role: user.Role,
	})
}

func (uh *UserHandler) GetAllUsers(c *gin.Context) {
	// get all users
	users, err := uh.svc.GetAllUsers(c)
	if err != nil {
		util.ResponseHandler(c, http.StatusBadRequest, true, err.Error())
		return
	}

	// response
	c.JSON(http.StatusOK, users)
}
