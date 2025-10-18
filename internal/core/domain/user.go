package domain

import "gorm.io/gorm"

type Role uint16

const (
	AdminRole Role = 5150
	TenantRole Role = 1984
	UserRole Role = 2001
)

type User struct {
	*gorm.Model

	Name string `json:"name" gorm:"size:255;not null"`
	Email string `json:"email" gorm:"size:255;not null;unique"`
	Provider string `json:"provider" gorm:"size:255;not null"`
	Password string `json:"password" gorm:"size:255"`
	Role Role `json:"role" gorm:"default:2001"`
	Picture string `json:"picture" gorm:"size:255"`
}
