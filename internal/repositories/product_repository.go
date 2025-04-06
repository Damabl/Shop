package repositories

import (
	"Shop/internal/models"
	"errors"
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
func (r *ProductRepository) GetProductByName(name string) (*models.Product, error) {
	var product models.Product
	err := r.Db.Where("name = ?", name).First(&product).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Это не ошибка — просто продукта нет
		return nil, nil
	}
	if err != nil {
		// Другая ошибка, например, с подключением к БД
		return nil, err
	}
	return &product, nil
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
