package config

import (
	"os"

	"github.com/joho/godotenv"
)

type (
	Container struct {
		HTTP  *HTTP
		DB    *DB
		Token *Token
	}

	HTTP struct {
		Host           string
		Port           string
		AllowedOrigins string
		AppEnv string
	}

	DB struct {
		Name     string
		User     string
		Password string
		Host     string
		Port     string
	}

	Token struct {
		Secret   string
		Duration string
	}
)

func Init() (*Container, error) {
	err := godotenv.Load()
	if err != nil {
		return &Container{}, err
	}

	http := &HTTP{
		Host: os.Getenv("HTTP_HOST"),
		Port: os.Getenv("HTTP_PORT"),
		AllowedOrigins: os.Getenv("HTTP_ALLOWED_ORIGINS"),
		AppEnv: os.Getenv("APP_ENV"),
	}

	db := &DB{
		Name: os.Getenv("DB_NAME"),
		User: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
	}

	token := &Token{
		Secret: os.Getenv("JWT_SECRET"),
		Duration: os.Getenv("TOKEN_DURATION"),
	}

	return &Container{
		HTTP: http,
		DB: db,
		Token: token,
	}, nil
}
