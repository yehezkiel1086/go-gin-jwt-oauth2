package handler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/adapter/config"
	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/core/domain"
	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/core/port"
	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/core/util"
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

	// generate jwt
	ss, duration, err := util.GenerateJWT(ah.jwtConf, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
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

func (ah *AuthHandler) GoogleLogin(c *gin.Context) {
	url := ah.svc.GetGoogleLoginURL(c)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (ah *AuthHandler) GoogleCallback(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing code"})
		return
	}

	ss, duration, err := ah.svc.GoogleCallback(c, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// set cookie
	c.SetCookie("jwt_token", ss, duration * 60, "/", ah.httpConf.Host, os.Getenv("APP_ENV") == "production", true)

	redirect := os.Getenv("GOOGLE_CLIENT_REDIRECT_URL")
	c.Redirect(http.StatusTemporaryRedirect, redirect)
}

func (ah *AuthHandler) Logout(c *gin.Context) {
	// remove jwt_token cookie
	c.SetCookie("jwt_token", "", -1, "/", ah.httpConf.Host, os.Getenv("APP_ENV") == "production", true)

	// return http response
	c.JSON(http.StatusOK, gin.H{
		"message": "logged out successfully",
	})
}
