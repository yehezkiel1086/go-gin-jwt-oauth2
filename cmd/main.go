package main

import (
	"context"
	"fmt"

	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/adapter/config"
	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/adapter/handler"
	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/adapter/storage/postgres"
	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/adapter/storage/postgres/repository"
	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/core/domain"
	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/core/service"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func main() {
	// get .env configs
	conf, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("✅ .env configs imported successfully")

	ctx := context.Background()
	
	// init db
	db, err := postgres.InitDB(ctx, conf.DB)
	if err != nil {
		panic(err)
	}
	fmt.Println("✅ postgres db initialized successfully")

	// migrate dbs
	if err := db.MigrateDB(&domain.User{}, &domain.Employee{}); err != nil {
		panic(err)
	}
	fmt.Println("✅ dbs migrated successfully")

	// dependency injection
	userRepo := repository.InitUserRepository(db)
	userSvc := service.InitUserService(userRepo)
	userHandler := handler.InitUserHandler(userSvc)

	authSvc := service.InitAuthService(userRepo, &oauth2.Config{
		ClientID:     conf.OAuth.ClientID,
		ClientSecret: conf.OAuth.ClientSecret,
		RedirectURL:  conf.OAuth.RedirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}, conf.JWT)
	authHandler := handler.InitAuthHandler(authSvc, conf.JWT, conf.HTTP)

	empRepo := repository.InitEmployeeRepository(db)
	empSvc := service.InitEmployeeService(empRepo)
	empHandler := handler.InitEmployeeHandler(empSvc)

	// routing
	r := handler.InitRouter(
		conf.HTTP,
		*userHandler,
		*authHandler,
		*empHandler,
	)
	fmt.Println("✅ routes initialized successfully")

	// run server
	if err := r.Start(); err != nil {
		panic(err)
	}
}
