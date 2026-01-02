package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yehezkiel1086/go-gin-nextjs-auth/internal/adapter/config"
)

type Router struct {
	r *gin.Engine
}

func NewRouter(
	userHandler *UserHandler,
) (*Router) {
	r := gin.New()

	// group routes
	pb := r.Group("/api/v1")

	// public user routes
	pb.POST("/register", userHandler.RegisterUser)

	// admin user routes
	pb.GET("/users", userHandler.GetUsers)

	return &Router{r}
}

func (r *Router) Serve(conf *config.HTTP) error {
	uri := conf.Host + ":" + conf.Port
	return r.r.Run(uri)
}
