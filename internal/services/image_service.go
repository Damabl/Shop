package services

import (
	"Shop/internal/cloud"
	"Shop/internal/db"
	"Shop/internal/models"
	"fmt"
	"mime/multipart"
)

type ImageService struct {
	CloudinaryService *cloud.CloudinaryService
}

func NewImageService(cloudinaryService *cloud.CloudinaryService) *ImageService {
	return &ImageService{CloudinaryService: cloudinaryService}
}

func (s *ImageService) UploadImage(file multipart.File, fileHeader *multipart.FileHeader, productID uint) (*models.Image, error, error) {
	imageURL, publicID, err := s.CloudinaryService.UploadImage(file, fileHeader)
	if err != nil {
		return nil, fmt.Errorf("failed to upload image: %w", err), nil
	}

	image := &models.Image{
		URL:       imageURL,
		PublicID:  publicID,
		ProductID: productID,
	}

	result := db.DB.Create(image)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to save image record in database: %w", result.Error), nil
	}

	return image, nil, nil
}

func (s *ImageService) DeleteImage(imageID uint) error {
	var image models.Image
	if err := db.DB.First(&image, imageID).Error; err != nil {
		return fmt.Errorf("image not found: %w", err)
	}

	err := s.CloudinaryService.DeleteImage(image.PublicID)
	if err != nil {
		return fmt.Errorf("failed to delete image from Cloudinary: %w", err)
	}

	if err := db.DB.Delete(&image).Error; err != nil {
		return fmt.Errorf("failed to delete image record from database: %w", err)
	}

	return nil
}
