package domain

import (
	"time"

	"gorm.io/gorm"
)

type Role uint16

const (
	AdminRole Role = 5150
	UserRole  Role = 2001
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// credentials
	Username string `json:"username" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null"`

	// information details
	Fullname string `json:"fullname" gorm:"not null"`
	Email    string `json:"email" gorm:"not null;unique"`
	Role Role `json:"role" gorm:"not null; default: 2001"`
}
