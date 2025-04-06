package models

import "gorm.io/gorm"

type User struct {
	gorm.Model        // Содержит id, created_at, updated_at, deleted_at
	Email      string `json:"email" gorm:"unique"`
	Name       string `json:"name"`
	Gender     string `json:"gender"`
	Age        int    `json:"age" gorm:"default:18;check:age >= 18 AND age <= 100"`
	Password   string `json:"-"`
	Role       string `json:"role"`
}
