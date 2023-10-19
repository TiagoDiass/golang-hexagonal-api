package usecase

import (
	"testing"

	"github.com/TiagoDiass/golang-hexagonal-api/internal/entity"
	"github.com/TiagoDiass/golang-hexagonal-api/internal/test"
)

func TestShouldDeleteProductSuccessfully(t *testing.T) {
	t.Parallel()

	productRepository := test.NewInMemoryProductRepository()
	deleteProductUsecase := NewDeleteProductUsecase(productRepository)

	productRepository.Create(&entity.Product{
		ID:    "fake-product-id",
		Name:  "fake-product",
		Price: 100,
	})

	err := deleteProductUsecase.Execute(DeleteProductInputDTO{
		ID: "fake-product-id",
	})

	if err != nil {
		t.Error("Expected not to receive an error from DeleteProductUsecase")
	}
}

func TestShouldReturnErrorWhenProductDoesNotExistCase(t *testing.T) {
	t.Parallel()

	productRepository := test.NewInMemoryProductRepository()
	deleteProductUsecase := NewDeleteProductUsecase(productRepository)

	err := deleteProductUsecase.Execute(DeleteProductInputDTO{
		ID: "fake-product-id",
	})

	if err == nil {
		t.Error("Expected to receive an error from DeleteProductUsecase")
	}
}
