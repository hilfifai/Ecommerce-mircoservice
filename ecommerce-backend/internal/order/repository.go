package order

import (
	"ecommerce-backend/internal/domain"

	"gorm.io/gorm"
)

type Repository interface {
	FindOrderByID(id uint) (*domain.Order, error)
	FindOrdersByUserID(userID uint) ([]domain.Order, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindOrderByID(id uint) (*domain.Order, error) {
	var order domain.Order
	err := r.db.Preload("OrderItems.Product").First(&order, id).Error
	return &order, err
}
func (r *repository) FindOrdersByUserID(userID uint) ([]domain.Order, error) {
	var orders []domain.Order
	err := r.db.Where("user_id = ?", userID).Preload("OrderItems.Product").Find(&orders).Error
	return orders, err
}