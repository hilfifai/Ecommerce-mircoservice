package database

import (
	"ecommerce-backend/internal/domain"
	"log"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	err := db.AutoMigrate(
		&domain.User{},
		&domain.Product{},
		&domain.Order{},
		&domain.OrderItem{},
	)
	if err != nil {
		log.Fatalf("Could not migrate database: %v", err)
	}
	log.Println("Database migration completed successfully.")
}