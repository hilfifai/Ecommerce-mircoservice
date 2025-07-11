package order

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
)

type Service interface {
	CreateOrder(orderRequest Request) error
}

type service struct {
	repo   Repository
	redis  *redis.Client
}

func NewService(repo Repository, redis *redis.Client) Service {
	return &service{repo, redis}
}

func (s *service) CreateOrder(orderRequest Request) error {
	orderJSON, err := json.Marshal(orderRequest)
	if err != nil {
		return err
	}
	
	return s.redis.LPush(context.Background(), "new_orders", orderJSON).Err()
}