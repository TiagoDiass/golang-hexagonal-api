package usecase

import (
	"testing"

	"github.com/TiagoDiass/golang-hexagonal-api/internal/entity"
	"github.com/TiagoDiass/golang-hexagonal-api/internal/test"
	"github.com/TiagoDiass/golang-hexagonal-api/internal/test/assertions"
)

func TestShouldFindProductByIdSuccessfully(t *testing.T) {
	t.Parallel()

	productRepository := test.NewInMemoryProductRepository()
	findProductByIdUsecase := NewFindProductByIdUsecase(productRepository)

	productRepository.Create(&entity.Product{
		ID:    "fake-product-id",
		Name:  "fake-product",
		Price: 100,
	})

	product, err := findProductByIdUsecase.Execute(FindProductByIdInputDTO{
		ID: "fake-product-id",
	})

	if err != nil {
		t.Error("Expected not to receive an error from FindProductByIdUsecase")
	}

	assertions.ExpectToBe(t, product.ID, "fake-product-id")
	assertions.ExpectToBe(t, product.Name, "fake-product")
}

func TestShouldReturnErrorWhenProductDoesNotExist(t *testing.T) {
	t.Parallel()

	productRepository := test.NewInMemoryProductRepository()
	findProductByIdUsecase := NewFindProductByIdUsecase(productRepository)

	product, err := findProductByIdUsecase.Execute(FindProductByIdInputDTO{
		ID: "fake-product-id",
	})

	if err == nil {
		t.Error("Expected to receive an error from FindProductByIdUsecase")
		assertions.ExpectToBe(t, err.Error(), "product does not exist")
	}

	if product != nil {
		t.Error("Expected to receive a null product from FindProductByIdUsecase")
	}
}
