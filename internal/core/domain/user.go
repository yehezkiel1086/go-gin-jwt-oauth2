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

	Email string `json:"email" gorm:"size:255;not null;unique"`
	Password string `json:"password" gorm:"size:255;not null"`
	Role Role `json:"role" gorm:"not null;default:2001"`
}
