package services

import (
	"Shop/internal/models"
	"Shop/internal/repositories"
	"errors"
	"gorm.io/gorm"
)

type ProductService struct {
	Repo         *repositories.ProductRepository
	ImageService *ImageService
}

func NewProductService(repo *repositories.ProductRepository, service *ImageService) *ProductService {
	return &ProductService{Repo: repo, ImageService: service}
}

func (s *ProductService) CreateProduct(product models.Product) error {
	existingProduct, err := s.Repo.GetProductByName(product.Name)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err // вернём ошибку, если это не "not found"
	}
	if existingProduct != nil {
		return errors.New("product already exists")
	}
	return s.Repo.CreateProduct(&product)
}

func (s *ProductService) GetProductByName(name string) (*models.Product, error) {
	return s.Repo.GetProductByName(name)
}
func (s *ProductService) UpdateProduct(product *models.Product) error {
	return s.Repo.UpdateProduct(&product)
}

func (s *ProductService) DeleteProduct(productID uint) error {
	product, err := s.Repo.GetProductById(productID)
	if err != nil {
		return err
	}
	errImage := s.ImageService.DeleteImage(product.ID)
	if errImage != nil {

		return err
	}
	return s.Repo.DeleteProduct(product)
}

func (s *ProductService) GetProductByID(id uint) (*models.Product, error) {
	return s.Repo.GetProductById(id)
}

func (s *ProductService) GetAllProducts(page, limit int) ([]*models.Product, error) {
	return s.Repo.GetAllProducts(page, limit)
}
