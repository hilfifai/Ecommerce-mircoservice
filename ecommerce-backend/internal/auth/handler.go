package auth

import (
	"ecommerce-backend/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(router *gin.RouterGroup, service Service) {
	h := &Handler{service}
	authRoutes := router.Group("/auth")
	authRoutes.POST("/register", h.Register)
	authRoutes.POST("/login", h.Login)
	authRoutes.OPTIONS("/*any", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})
}

func (h *Handler) Register(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Register(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (h *Handler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}