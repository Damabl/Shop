package repositories

import (
	"Shop/internal/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	Db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{Db: db}
}
func (repo *ProductRepository) CreateProduct(product *models.Product) error {
	return repo.Db.Create(product).Error
}
func (repo *ProductRepository) UpdateProduct(product *models.Product) error {
	return repo.Db.Save(product).Error
}
func (repo *ProductRepository) DeleteProduct(product *models.Product) error {
	return repo.Db.Delete(product).Error
}
func (repo *ProductRepository) GetProductByName(name string) (*models.Product, error) {
	product := &models.Product{}
	err := repo.Db.Where("name = ?", name).First(product).Error
	return product, err
}
func (repo *ProductRepository) GetProductById(id uint) (*models.Product, error) {
	product := &models.Product{}
	if err := repo.Db.Where("id = ?", id).First(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}
func (repo *ProductRepository) GetAllProducts(page, limit int) ([]*models.Product, error) {
	var products []*models.Product
	offset := (page - 1) * limit

	err := repo.Db.Limit(limit).Offset(offset).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
