package services

import (
	"Shop/internal/models"
	"Shop/internal/repositories"
	"errors"
)

type CartService struct {
	Repo *repositories.CartRepository
}

func NewCartService(repo *repositories.CartRepository) *CartService {
	return &CartService{Repo: repo}
}

// Получение корзины
func (s *CartService) GetCart(userID uint) (*models.Cart, error) {
	return s.Repo.GetCartByUserID(userID)
}

// Добавление товара в корзину
func (s *CartService) AddToCart(userID, productID uint, quantity int) error {
	if quantity <= 0 {
		return errors.New("quantity must be greater than zero")
	}

	cartItem := &models.CartItem{
		CartID:    userID, // CartID = UserID (предполагаем, что у каждого юзера одна корзина)
		ProductID: productID,
		Quantity:  quantity,
	}

	return s.Repo.AddItem(cartItem)
}

// Обновление количества товара
func (s *CartService) UpdateCartItem(cartItemID uint, quantity int) error {
	if quantity <= 0 {
		return errors.New("quantity must be greater than zero")
	}

	cartItem := &models.CartItem{
		ID:       cartItemID,
		Quantity: quantity,
	}

	return s.Repo.UpdateItemQuantity(cartItem)
}

// Удаление товара из корзины
func (s *CartService) RemoveFromCart(cartItemID uint) error {
	return s.Repo.RemoveItem(cartItemID)
}

// Очистка корзины
func (s *CartService) ClearCart(userID uint) error {
	return s.Repo.ClearCart(userID)
}
