package product

import (
	"bytes"
	"ecommerce-backend/internal/domain"
	"encoding/json"
	"net/http"
)

type Service interface {
	CreateProduct(product *domain.Product) error
}

type service struct {
	repo          Repository
	n8nWebhookURL string
}

func NewService(repo Repository, n8nWebhookURL string) Service {
	return &service{repo, n8nWebhookURL}
}

func (s *service) CreateProduct(product *domain.Product) error {
	if err := s.repo.Create(product); err != nil {
		return err
	}

	go s.notifyN8n(product)
	return nil
}

func (s *service) notifyN8n(product *domain.Product) {
	payload, err := json.Marshal(product)
	if err != nil {
		return
	}

	http.Post(s.n8nWebhookURL, "application/json", bytes.NewBuffer(payload))
}