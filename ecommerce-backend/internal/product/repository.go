package product

import (
	"ecommerce-backend/internal/domain"

	"gorm.io/gorm"
)

type Repository interface {
	Create(product *domain.Product) error
	FindByID(id uint) (*domain.Product, error)
	FindAll() ([]domain.Product, error)
	Update(product *domain.Product) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(product *domain.Product) error {
	return r.db.Create(product).Error
}

func (r *repository) FindByID(id uint) (*domain.Product, error) {
	var product domain.Product
	err := r.db.First(&product, id).Error
	return &product, err
}

func (r *repository) FindAll() ([]domain.Product, error) {
	var products []domain.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *repository) Update(product *domain.Product) error {
	return r.db.Save(product).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&domain.Product{}, id).Error
}