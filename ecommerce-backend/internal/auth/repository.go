package auth

import (
	"ecommerce-backend/internal/domain"
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user *domain.User) error
	FindUserByEmail(email string) (*domain.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) CreateUser(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *repository) FindUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}