package postgres

import (
	"context"
	"fmt"

	"github.com/yehezkiel1086/go-gin-jwt-oauth2/internal/adapter/config"
	"github.com/yehezkiel1086/go-gin-jwt-oauth2/internal/core/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

func ConnectDB(ctx context.Context, conf *config.DB) (*DB, error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta", conf.Host, conf.User, conf.Password, conf.Name, conf.Port)
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return &DB{}, err
	}

	return &DB{db: db}, nil
}

func (db *DB) Migrate() error {
	err := db.db.AutoMigrate(&domain.User{})
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) GetDB() *gorm.DB {
	return db.db
}
