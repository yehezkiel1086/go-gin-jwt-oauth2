package main

import (
	"fmt"

	"github.com/yehezkiel1086/go-gin-jwt-oauth2/internal/adapter/config"
)

func main() {
	// get .env config
	conf, err := config.Init()
	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta", conf.DB.Host, conf.DB.User, conf.DB.Password, conf.DB.Name, conf.DB.Port)

	fmt.Println(dsn)

	// connect db
}
