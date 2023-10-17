package usecase

import (
	"errors"

	"github.com/TiagoDiass/golang-hexagonal-api/internal/repository"
)

type DeleteProductInputDTO struct {
	ID string
}

type DeleteProductUsecase struct {
	ProductRepository repository.ProductRepository
}

func NewDeleteProductUsecase(productRepository repository.ProductRepository) *DeleteProductUsecase {
	return &DeleteProductUsecase{
		ProductRepository: productRepository,
	}
}

func (u *DeleteProductUsecase) Execute(input DeleteProductInputDTO) error {
	product, err := u.ProductRepository.FindById(input.ID)

	if err != nil {
		return err
	}

	if product == nil {
		errorMessage := "product does not exist"

		return errors.New(errorMessage)
	}

	deleteErr := u.ProductRepository.Delete(input.ID)

	if deleteErr != nil {
		return deleteErr
	}

	return nil
}
