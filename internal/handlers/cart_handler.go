package handlers

import (
	"Shop/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CartHandler struct {
	Service *services.CartService
}

func NewCartHandler(service *services.CartService) *CartHandler {
	return &CartHandler{Service: service}
}

// Получение корзины пользователя
func (h *CartHandler) GetCart(ctx *gin.Context) {
	userID, _ := strconv.Atoi(ctx.Param("user_id"))
	cart, err := h.Service.GetCart(uint(userID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}
	ctx.JSON(http.StatusOK, cart)
}

// Добавление товара в корзину
func (h *CartHandler) AddToCart(ctx *gin.Context) {
	userID, _ := strconv.Atoi(ctx.Param("user_id"))
	productID, _ := strconv.Atoi(ctx.Param("product_id"))
	quantity, _ := strconv.Atoi(ctx.Query("quantity"))

	if err := h.Service.AddToCart(uint(userID), uint(productID), quantity); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Product added to cart"})
}

// Обновление количества товара
func (h *CartHandler) UpdateCartItem(ctx *gin.Context) {
	cartItemID, _ := strconv.Atoi(ctx.Param("cart_item_id"))
	quantity, _ := strconv.Atoi(ctx.Query("quantity"))

	if err := h.Service.UpdateCartItem(uint(cartItemID), quantity); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Cart item updated"})
}

// Удаление товара из корзины
func (h *CartHandler) RemoveFromCart(ctx *gin.Context) {
	cartItemID, _ := strconv.Atoi(ctx.Param("cart_item_id"))

	if err := h.Service.RemoveFromCart(uint(cartItemID)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove item"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Item removed from cart"})
}

// Очистка корзины
func (h *CartHandler) ClearCart(ctx *gin.Context) {
	userID, _ := strconv.Atoi(ctx.Param("user_id"))

	if err := h.Service.ClearCart(uint(userID)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to clear cart"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Cart cleared"})
}
