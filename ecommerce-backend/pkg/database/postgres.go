package database

import (
	"ecommerce-backend/config"
	"fmt"
	"log"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.DBSslMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connection established")
	return db
}

func parseRedisDB(dbStr string) int {
	db, err := strconv.Atoi(dbStr)
	if err != nil {
		return 0 // Default to DB 0
	}
	return db
}