package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
	Quantity    int     `json:"quantity"`
	Brand       string  `json:"brand"`
	Category    string  `json:"category"`
	Size        string  `json:"size"`
	Color       string  `json:"color"`
	Material    string  `json:"material"`
	Gender      string  `json:"gender"`
	Season      string  `json:"season"`
	Discount    float64 `json:"discount"`
}
