package cloud

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type CloudinaryService struct {
	cld *cloudinary.Cloudinary
}

func NewCloudinaryService() (*CloudinaryService, error) {
	cld, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	if err != nil {
		return nil, fmt.Errorf("failed to init Cloudinary: %w", err)
	}

	return &CloudinaryService{cld: cld}, nil
}

func (s *CloudinaryService) UploadImage(file multipart.File, fileHeader *multipart.FileHeader) (string, string, error) {
	ctx := context.Background()

	extension := filepath.Ext(fileHeader.Filename)
	uniqueName := fmt.Sprintf("%s%s", uuid.New().String(), extension)

	uploadResult, err := s.cld.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID: uniqueName,
		Folder:   "products",
	})
	if err != nil {
		log.Println("Cloudinary upload error:", err)
		return "", "", err
	}
	return uploadResult.SecureURL, uniqueName, nil
}
func (s *CloudinaryService) DeleteImage(publicID string) error {
	ctx := context.Background()
	_, err := s.cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: publicID,
	})
	if err != nil {
		log.Println("Cloudinary delete error:", err)
		return fmt.Errorf("failed to delete image: %w", err)
	}
	return nil
}
