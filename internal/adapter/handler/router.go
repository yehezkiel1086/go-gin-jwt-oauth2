package handler

import (
	"fmt"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/adapter/config"
	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/core/domain"
)

type Router struct {
	r *gin.Engine
	conf *config.HTTP
}

func InitRouter(
	conf *config.HTTP,
	userHandler UserHandler,
	authHandler AuthHandler,
	empHandler EmployeeHandler,
) *Router {
	r := gin.New()

	// --- CORS middleware ---
	r.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Split(conf.AllowedOrigins, ","),
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	// route groupings (public, user and admin)
	pb := r.Group("/api/v1")
	us := pb.Group("/").Use(AuthMiddleware())
	ad := pb.Group("/admin").Use(AuthMiddleware(), RoleMiddleware(domain.AdminRole))

	// auth routes
	// jwt auth (public)
	pb.POST("/register", userHandler.Register)
	pb.POST("/login", authHandler.Login)

	// google oauth (public)
	pb.GET("/login/google", authHandler.GoogleLogin)
	pb.GET("/login/google/callback", authHandler.GoogleCallback)

	// auth (public)
	us.POST("/logout", authHandler.Logout)

	// employee routes
	// user
	us.GET("/employees", empHandler.GetEmployees)

	// admin
	ad.POST("/employees", empHandler.CreateEmployee)

	return &Router{r: r, conf: conf}
}

func (r *Router) Start() error {
	uri := fmt.Sprintf("%v:%v", r.conf.Host, r.conf.Port)
	if err := r.r.Run(uri); err != nil {
		return err
	}

	return nil
}
