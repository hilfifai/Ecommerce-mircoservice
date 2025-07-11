package product

import (
	"ecommerce-backend/internal/auth"
	"ecommerce-backend/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service     Service
	authService auth.Service
}

func NewHandler(router *gin.RouterGroup, service Service, authService auth.Service) {
	h := &Handler{service, authService}

	productRoutes := router.Group("/products")
	// Di masa depan, Anda mungkin ingin melindungi rute ini juga
	productRoutes.POST("", h.CreateProduct)
}

func (h *Handler) CreateProduct(c *gin.Context) {
	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, product)
}