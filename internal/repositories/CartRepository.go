package repositories

import (
	"Shop/internal/models"
	"gorm.io/gorm"
)

type CartRepository struct {
	Db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{Db: db}
}

// Получение корзины пользователя
func (repo *CartRepository) GetCartByUserID(userID uint) (*models.Cart, error) {
	var cart models.Cart
	err := repo.Db.Preload("Items.Product").Where("user_id = ?", userID).First(&cart).Error
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

// Добавление товара в корзину
func (repo *CartRepository) AddItem(cartItem *models.CartItem) error {
	return repo.Db.Create(cartItem).Error
}

// Обновление количества товара в корзине
func (repo *CartRepository) UpdateItemQuantity(cartItem *models.CartItem) error {
	return repo.Db.Save(cartItem).Error
}

// Удаление товара из корзины
func (repo *CartRepository) RemoveItem(cartItemID uint) error {
	return repo.Db.Delete(&models.CartItem{}, cartItemID).Error
}

// Очистка корзины
func (repo *CartRepository) ClearCart(userID uint) error {
	return repo.Db.Where("cart_id IN (SELECT id FROM carts WHERE user_id = ?)", userID).Delete(&models.CartItem{}).Error
}
