package domain

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID      uint        `json:"user_id"`
	User        User        `gorm:"foreignKey:UserID"`
	TotalAmount float64     `json:"total_amount"`
	Status      string      `json:"status"` // e.g., PENDING, PROCESSED, FAILED
	OrderItems  []OrderItem `json:"order_items"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}