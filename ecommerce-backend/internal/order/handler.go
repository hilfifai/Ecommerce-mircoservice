package order

import (
	"ecommerce-backend/internal/auth"
	"net/http"
	"ecommerce-backend/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Handler struct {
	service     Service
	authService auth.Service
}

type Request struct {
	UserID uint `json:"user_id"`
	Items  []struct {
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	} `json:"items"`
}

func NewHandler(router *gin.RouterGroup, service Service, authService auth.Service) {
	h := &Handler{service, authService}

	orderRoutes := router.Group("/orders")
	orderRoutes.Use(middleware.Middleware(authService))
	orderRoutes.POST("", h.CreateOrder)
}

func (h *Handler) CreateOrder(c *gin.Context) {
	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["user_id"].(float64))
	req.UserID = userID

	if err := h.service.CreateOrder(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to place order"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Order received and is being processed"})
}