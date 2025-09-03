package main

import (
	"context"
	"fmt"

	"github.com/yehezkiel1086/go-gin-jwt-oauth2/internal/adapter/config"
	"github.com/yehezkiel1086/go-gin-jwt-oauth2/internal/adapter/storage/postgres"
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
}
