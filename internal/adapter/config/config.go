package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type (
	Container struct {
		App  *App
		HTTP *HTTP
		DB   *DB
		JWT  *JWT
		OAuth *OAuth
	}

	App struct {
		Name string
		Env  string
	}

	HTTP struct {
		Host string
		Port string
		AllowedOrigins string
	}

	DB struct {
		Name string
		User string
		Pass string
		Host string
		Port string
	}

	OAuth struct {
		ClientID string
		ClientSecret string
		RedirectURL string
	}

	JWT struct {
		Secret string
		Duration string
	}
)

func InitConfig() (*Container, error) {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			return &Container{}, errors.New("error loading .env file")
		}
	}

	App := &App{
		Name: os.Getenv("APP_NAME"),
		Env: os.Getenv("APP_ENV"),
	}

	HTTP := &HTTP{
		Host: os.Getenv("HTTP_HOST"),
		Port: os.Getenv("HTTP_PORT"),
		AllowedOrigins: os.Getenv("HTTP_ALLOWED_ORIGINS"),
	}

	DB := &DB{
		Name: os.Getenv("DB_NAME"),
		User: os.Getenv("DB_USER"),
		Pass: os.Getenv("DB_PASS"),
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
	}

	OAuth := &OAuth{
		ClientID: os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL: os.Getenv("GOOGLE_REDIRECT_URL"),
	}

	JWT := &JWT{
		Secret: os.Getenv("JWT_SECRET"),
		Duration: os.Getenv("TOKEN_DURATION"),
	}
	
	return &Container{
		App: App,
		HTTP: HTTP,
		DB: DB,
		OAuth: OAuth,
		JWT: JWT,
	}, nil
}
