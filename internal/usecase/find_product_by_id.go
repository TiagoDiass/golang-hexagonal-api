package usecase

import (
	"errors"

	"github.com/TiagoDiass/golang-hexagonal-api/internal/repository"
)

type FindProductByIdInputDTO struct {
	ID string
}

type FindProductByIdOutputDTO struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type FindProductByIdUsecase struct {
	ProductRepository repository.ProductRepository
}

func NewFindProductByIdUsecase(productRepository repository.ProductRepository) *FindProductByIdUsecase {
	return &FindProductByIdUsecase{
		ProductRepository: productRepository,
	}
}

func (u *FindProductByIdUsecase) Execute(input FindProductByIdInputDTO) (*FindProductByIdOutputDTO, error) {
	product, err := u.ProductRepository.FindById(input.ID)

	if err != nil {
		return nil, err
	}

	if product == nil {
		return nil, errors.New("product does not exist")
	}

	output := &FindProductByIdOutputDTO{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}

	return output, nil
}
