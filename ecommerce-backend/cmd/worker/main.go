package main

import (
	"ecommerce-backend/config"
	"ecommerce-backend/internal/worker"
	"ecommerce-backend/pkg/database"
	"ecommerce-backend/pkg/queue"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	db := database.InitDB(cfg)
	redisClient := queue.InitRedis(cfg)

	processor := worker.NewOrderProcessor(db, redisClient)

	log.Println("Starting order processing worker...")
	go processor.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down worker...")
}