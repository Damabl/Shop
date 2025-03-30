package services

import (
	"Shop/internal/models"
	"Shop/internal/repositories"
	"errors"
)

type ProductService struct {
	Repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{Repo: repo}
}

func (s *ProductService) CreateProduct(product models.Product) error {
	existingProduct, _ := s.Repo.GetProductByName(product.Name)
	if existingProduct != nil {
		return errors.New("product already exists")
	}

	return s.Repo.CreateProduct(&product)
}
func (s *ProductService) GetProductByName(name string) (*models.Product, error) {
	return s.Repo.GetProductByName(name)
}
func (s *ProductService) UpdateProduct(product models.Product) error {
	return s.Repo.UpdateProduct(&product)
}

func (s *ProductService) DeleteProduct(productID uint) error {
	product, err := s.Repo.GetProductById(productID)
	if err != nil {
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
