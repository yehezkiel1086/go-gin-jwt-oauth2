package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func InitRouter(
	uh UserHandler,
) *Router {
	r := gin.New()

	// public routes
	v1 := r.Group("/api/v1") 
	v1.POST("/register", uh.Register)

	return &Router{r}
}

func (r *Router) Serve(addr string) {
	fmt.Println("Server is running on", addr)
	r.Run(addr)
}
