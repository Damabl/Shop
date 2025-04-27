package handlers

import (
	"Shop/internal/cloud"
	"Shop/internal/models"
	"Shop/internal/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	Service           *services.ProductService
	CloudinaryService *cloud.CloudinaryService
	ImageService      *services.ImageService
}

func NewProductHandler(service *services.ProductService, cloudinary *cloud.CloudinaryService, image *services.ImageService) *ProductHandler {
	return &ProductHandler{Service: service, CloudinaryService: cloudinary, ImageService: image}
}

func (h *ProductHandler) CreateProduct(ctx *gin.Context) {
	var product models.Product

	if err := ctx.ShouldBind(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
		return
	}

	if err := h.Service.CreateProduct(product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) ImageUpload(ctx *gin.Context) {
	fileHeader, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Image is required"})
		return
	}

	productIdStr, ok := ctx.GetPostForm("productId")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
		return
	}

	productIdUint, err := strconv.ParseUint(productIdStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Product ID"})
		return
	}

	product, err := h.Service.GetProductByID(uint(productIdUint))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Product not found: %v", err)})
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot open image"})
		return
	}
	defer func(file multipart.File) {
		if err := file.Close(); err != nil {
			log.Println("Error closing file:", err)
		}
	}(file)

	image, _, err := h.ImageService.UploadImage(file, fileHeader, product.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image"})
		return
	}
	product.Images = append(product.Images, *image)
	if err := h.Service.UpdateProduct(product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to update product: %v", err)})
		return
	}

	return
}

func (h *ProductHandler) GetProductByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := h.Service.GetProductByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (h *ProductHandler) GetAllProducts(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	products, err := h.Service.GetAllProducts(page, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (h *ProductHandler) UpdateProduct(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var product models.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.ID = uint(id)
	err = h.Service.UpdateProduct(&product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

func (h *ProductHandler) DeleteProduct(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	err = h.Service.DeleteProduct(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
