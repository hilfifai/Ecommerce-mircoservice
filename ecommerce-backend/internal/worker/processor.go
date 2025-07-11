package worker

import (
	"context"
	"ecommerce-backend/internal/domain"
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type OrderProcessor struct {
	db    *gorm.DB
	redis *redis.Client
}

type OrderRequest struct {
	UserID uint `json:"user_id"`
	Items  []struct {
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	} `json:"items"`
}

func NewOrderProcessor(db *gorm.DB, redis *redis.Client) *OrderProcessor {
	return &OrderProcessor{db, redis}
}

func (p *OrderProcessor) Start() {
	for {
		result, err := p.redis.BRPop(context.Background(), 0, "new_orders").Result()
		if err != nil {
			log.Printf("Error fetching order from queue: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		var req OrderRequest
		if err := json.Unmarshal([]byte(result[1]), &req); err != nil {
			log.Printf("Error unmarshalling order: %v", err)
			continue
		}

		p.process(req)
	}
}

func (p *OrderProcessor) process(req OrderRequest) {
	tx := p.db.Begin()
	if tx.Error != nil {
		log.Printf("Error starting transaction: %v", tx.Error)
		return
	}
	
	defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
        }
    }()

	var totalAmount float64
	var orderItems []domain.OrderItem

	for _, item := range req.Items {
		var product domain.Product
		if err := tx.First(&product, item.ProductID).Error; err != nil {
			log.Printf("Product not found: %d", item.ProductID)
			tx.Rollback()
			return
		}

		if product.Stock < item.Quantity {
			log.Printf("Insufficient stock for product %d", item.ProductID)
			tx.Rollback()
			return
		}

		newStock := product.Stock - item.Quantity
		if err := tx.Model(&product).Update("stock", newStock).Error; err != nil {
			log.Printf("Failed to update stock: %v", err)
			tx.Rollback()
			return
		}

		totalAmount += product.Price * float64(item.Quantity)
		orderItems = append(orderItems, domain.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     product.Price,
		})
	}

	order := domain.Order{
		UserID:      req.UserID,
		TotalAmount: totalAmount,
		Status:      "PROCESSED",
		OrderItems:  orderItems,
	}

	if err := tx.Create(&order).Error; err != nil {
		log.Printf("Failed to create order: %v", err)
		tx.Rollback()
		return
	}

	tx.Commit()
	log.Printf("Successfully processed order for user %d", req.UserID)
}