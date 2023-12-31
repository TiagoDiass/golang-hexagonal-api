package usecase

import (
	"github.com/TiagoDiass/golang-hexagonal-api/internal/entity"
	"github.com/TiagoDiass/golang-hexagonal-api/internal/repository"
)

type CreateProductInputDTO struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type CreateProductOutputDTO struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type CreateProductUsecase struct {
	ProductRepository repository.ProductRepository
}

func NewCreateProductUsecase(productRepository repository.ProductRepository) *CreateProductUsecase {
	return &CreateProductUsecase{
		ProductRepository: productRepository,
	}
}

func (u *CreateProductUsecase) Execute(input CreateProductInputDTO) (*CreateProductOutputDTO, error) {
	product := entity.NewProduct(input.Name, input.Price)

	err := u.ProductRepository.Create(product)

	if err != nil {
		return nil, err
	}

	return &CreateProductOutputDTO{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}, nil
}
