package repository

import "github.com/TiagoDiass/golang-hexagonal-api/internal/entity"

type ProductRepository interface {
	Create(product *entity.Product) error
	FindAll() ([]*entity.Product, error)
	FindById(id string) (*entity.Product, error)
	Delete(id string) error
}
