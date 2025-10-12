package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/adapter/config"
	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/core/domain"
)

type Router struct {
	*gin.Engine
}

func InitRouter(
	userHandler UserHandler,
	authHandler AuthHandler,
	empHandler EmployeeHandler,
) *Router {
	r := gin.New()

	// route groupings (public, user and admin)
	pb := r.Group("/api/v1")
	us := pb.Group("/").Use(AuthMiddleware())
	ad := pb.Group("/admin").Use(AuthMiddleware(), RoleMiddleware(domain.AdminRole))

	// auth routes
	// public
	pb.POST("/register", userHandler.Register)
	pb.POST("/login", authHandler.Login)

	// employee routes
	// user
	us.GET("/employees", empHandler.GetEmployees)

	// admin
	ad.POST("/employees", empHandler.CreateEmployee)

	return &Router{r}
}

func (r *Router) Start(conf *config.HTTP) error {
	uri := fmt.Sprintf("%v:%v", conf.Host, conf.Port)
	if err := r.Run(uri); err != nil {
		return err
	}

	return nil
}
