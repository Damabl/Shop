package models

type Image struct {
	ID        uint   `gorm:"primaryKey"`
	URL       string `json:"url"`
	PublicID  string `json:"public_id"` // для Cloudinary
	ProductID uint   `json:"product_id"`
}
