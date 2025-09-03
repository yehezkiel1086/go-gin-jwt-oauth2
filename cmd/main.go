package main

import (
	"context"
	"fmt"

	"github.com/yehezkiel1086/go-gin-jwt-oauth2/internal/adapter/config"
	"github.com/yehezkiel1086/go-gin-jwt-oauth2/internal/adapter/handler"
	"github.com/yehezkiel1086/go-gin-jwt-oauth2/internal/adapter/storage/postgres"
	"github.com/yehezkiel1086/go-gin-jwt-oauth2/internal/adapter/storage/postgres/repository"
	"github.com/yehezkiel1086/go-gin-jwt-oauth2/internal/core/service"
)

func main() {
	// get .env config
	conf, err := config.Init()
	if err != nil {
		panic(err)
	}

	// connect db
	ctx := context.Background()
	db, err := postgres.ConnectDB(ctx, conf.DB)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to db.")

	// migrate db
	if err := db.Migrate(); err != nil {
		panic(err)
	}
	fmt.Println("Migration success.")

	// dependency injection
	userRepository := repository.InitUserRepository(db)
	userService := service.InitUserService(userRepository)
	userHandler := handler.InitUserHandler(userService)

	// routing
	r := handler.InitRouter(*userHandler)

	// get address and serve
	addr := fmt.Sprintf("%v:%v", conf.HTTP.Host, conf.HTTP.Port)
	r.Serve(addr)
}
