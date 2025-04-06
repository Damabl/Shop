package models

import "gorm.io/gorm"

type CartItem struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	CartID    uint    `json:"cart_id"` // Связь с корзиной
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Product   Product `gorm:"foreignKey:ProductID"`
}

type Cart struct {
	gorm.Model
	UserID uint       `json:"user_id"`
	User   User       `gorm:"foreignKey:UserID"` // Добавьте эту связь
	Items  []CartItem `gorm:"foreignKey:CartID"`
}
