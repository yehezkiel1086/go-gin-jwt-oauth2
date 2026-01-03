package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yehezkiel1086/go-gin-nextjs-auth/internal/adapter/config"
	"github.com/yehezkiel1086/go-gin-nextjs-auth/internal/core/domain"
)

type Router struct {
	r *gin.Engine
}

func NewRouter(
	userHandler *UserHandler,
	authHandler *AuthHandler,
	jobHandler *JobHandler,
) (*Router) {
	r := gin.New()

	// group routes
	pb := r.Group("/api/v1")
	us := pb.Group("/", AuthMiddleware(), RoleMiddleware(domain.UserRole, domain.AdminRole))
	ad := pb.Group("/", AuthMiddleware(), RoleMiddleware(domain.AdminRole))

	// public user and auth routes
	pb.POST("/login", authHandler.Login)
	pb.POST("/register", userHandler.RegisterUser)

	// admin user routes
	ad.GET("/users", userHandler.GetUsers)

	// user job routes
	us.GET("/jobs", jobHandler.GetJobs)
	us.GET("/jobs/:id", jobHandler.GetJobById)

	// admin job routes
	ad.POST("/jobs", jobHandler.CreateJob)
	ad.DELETE("/jobs/:id", jobHandler.DeleteJob)

	return &Router{r}
}

func (r *Router) Serve(conf *config.HTTP) error {
	uri := conf.Host + ":" + conf.Port
	return r.r.Run(uri)
}
