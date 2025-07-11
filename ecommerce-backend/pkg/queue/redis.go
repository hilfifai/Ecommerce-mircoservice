package queue

import (
	"context"
	"ecommerce-backend/config"
	"log"
	"strconv"
  "os"
	"github.com/go-redis/redis/v8"
)

func InitRedis(cfg *config.Config) *redis.Client {
	redisDB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
		DB:       redisDB,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	log.Println("Redis connection established")
	return rdb
}