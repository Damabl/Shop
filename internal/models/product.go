package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string  `json:"name" form:"name"`
	Description string  `json:"description" form:"description"`
	Price       float64 `json:"price" form:"price"`
	Quantity    int     `json:"quantity" form:"quantity"`
	Brand       string  `json:"brand" form:"brand"`
	Category    string  `json:"category" form:"category"`
	Size        string  `json:"size" form:"size"`
	Color       string  `json:"color" form:"color"`
	Material    string  `json:"material" form:"material"`
	Gender      string  `json:"gender" form:"gender"`
	Season      string  `json:"season" form:"season"`
	Discount    float64 `json:"discount" form:"discount"`
	Image       string  `json:"image" form:"image"`
}
