package test

import "github.com/TiagoDiass/golang-hexagonal-api/internal/entity"

type InMemoryProductRepository struct {
	products map[string]*entity.Product
}

func NewInMemoryProductRepository() *InMemoryProductRepository {
	return &InMemoryProductRepository{
		products: make(map[string]*entity.Product),
	}
}

func (r *InMemoryProductRepository) Create(product *entity.Product) error {
	r.products[product.ID] = product

	return nil
}

func (r *InMemoryProductRepository) FindAll() ([]*entity.Product, error) {
	var products []*entity.Product

	for _, product := range r.products {
		products = append(products, product)
	}

	return products, nil
}

func (r *InMemoryProductRepository) FindById(productId string) (*entity.Product, error) {
	product, exists := r.products[productId]

	if !exists {
		return nil, nil
	}

	return product, nil
}

func (r *InMemoryProductRepository) Delete(productId string) error {
	delete(r.products, productId)
	return nil
}
