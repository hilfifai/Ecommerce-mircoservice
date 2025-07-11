package main

import (
	"ecommerce-backend/config"
	"ecommerce-backend/internal/auth"
	"ecommerce-backend/internal/order"
	"ecommerce-backend/internal/product"
	"ecommerce-backend/pkg/database"
	"ecommerce-backend/pkg/queue"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	db := database.InitDB(cfg)
	database.RunMigrations(db) 

	redisClient := queue.InitRedis(cfg)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{
			"http://localhost:8000",
			"http://localhost:8080",
			"http://localhost",
			"http://127.0.0.1:8000",
			"http://127.0.0.1:8080",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{
			"Origin",
			"Content-Type",
			"Authorization",
			"Accept",
			"X-Requested-With",
		},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	//router.Use(middleware.CORSMiddleware());
	api := router.Group("/api/v1")

	userRepo := auth.NewRepository(db)
	authService := auth.NewService(userRepo, cfg.JWTSecret)
	auth.NewHandler(api, authService)
	
	productRepo := product.NewRepository(db)
	productService := product.NewService(productRepo, cfg.N8nWebhookURL)
	product.NewHandler(api, productService, authService)

	orderRepo := order.NewRepository(db)
	orderService := order.NewService(orderRepo, redisClient)
	order.NewHandler(api, orderService, authService)

	log.Printf("Starting server on port %s", cfg.APIPort)
	router.Run(":" + cfg.APIPort)
}