package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/core/domain"
	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/core/port"
)

type UserHandler struct {
	svc port.UserService
}

func InitUserHandler(svc port.UserService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

type RegisterReq struct {
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (uh *UserHandler) Register(c *gin.Context) {
	// bind input
	var input *RegisterReq
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "name, email and password are required",
		})
		return
	}

	// register
	if _, err := uh.svc.Register(c, &domain.User{
		Name: input.Name,
		Email: input.Email,
		Password: input.Password,
	}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// return response
	c.JSON(http.StatusCreated, gin.H{
		"message": "user registered successfully",
	})
}
