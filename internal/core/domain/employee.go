package domain

import "gorm.io/gorm"

type Employee struct {
	gorm.Model

	Name string `json:"name" gorm:"size:255;unique;not null"`
	Position string `json:"position" gorm:"size:255;not null"`
	Description string `json:"description" gorm:"size:255"`
}
