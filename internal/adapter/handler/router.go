package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yehezkiel1086/go-gin-jwt-oauth2/internal/core/domain"
)

type Router struct {
	*gin.Engine
}

func InitRouter(
	uh UserHandler,
	ah AuthHandler,
) *Router {
	r := gin.New()

	// public routes
	v1 := r.Group("/api/v1")
	v1.POST("/register", uh.Register)
	v1.POST("/login", ah.Login)

	// user routes
	us := v1.Group("/")
	us.Use(AuthMiddleware())

	// admin routes
	ad := us.Group("/")
	ad.Use(RoleMiddleware(domain.AdminRole))

	// user routes
	ad.GET("/users", uh.GetAllUsers)
	us.GET("/users/:username", uh.GetUserByUsername)

	return &Router{r}
}

func (r *Router) Serve(addr string) {
	fmt.Println("Server is running on", addr)
	r.Run(addr)
}
