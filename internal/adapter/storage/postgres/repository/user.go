package repository

import (
	"context"
	"errors"

	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/adapter/storage/postgres"
	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/core/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *postgres.DB
}

func InitUserRepository(db *postgres.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	db := ur.db.GetDB()

	if err := db.Create(&user).Error; err != nil {
		return &domain.User{}, err
	}

	return user, nil
}

func (ur *UserRepository) CreateOrUpdate(ctx context.Context, user *domain.User) (*domain.User, error) {
    db := ur.db.GetDB()

    existing := &domain.User{}
    err := db.Where("email = ?", user.Email).First(existing).Error

    if errors.Is(err, gorm.ErrRecordNotFound) {
        // Create new user
        if err := db.Create(user).Error; err != nil {
            return nil, err
        }
        return user, nil
    } else if err != nil {
        // Some other DB error
        return nil, err
    }

    // Update existing user fields
    existing.Name = user.Name
    existing.Picture = user.Picture

    if err := db.Save(existing).Error; err != nil {
        return nil, err
    }

    return existing, nil
}

func (ur *UserRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	db := ur.db.GetDB()

	var user *domain.User
	if err := db.First(&user, "email = ?", email).Error; err != nil {
		return &domain.User{}, err
	}

	return user, nil
}
