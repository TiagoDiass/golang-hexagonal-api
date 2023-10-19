package usecase

import (
	"testing"

	"github.com/TiagoDiass/golang-hexagonal-api/internal/entity"
	"github.com/TiagoDiass/golang-hexagonal-api/internal/test"
	"github.com/TiagoDiass/golang-hexagonal-api/internal/test/assertions"
)

func TestShouldListAllProductsSuccessfully(t *testing.T) {
	t.Parallel()

	productRepository := test.NewInMemoryProductRepository()
	listProductsUsecase := NewListProductsUsecase(productRepository)

	productRepository.Create(&entity.Product{
		ID:    "first-product-id",
		Name:  "first-product-name",
		Price: 100,
	})

	productRepository.Create(&entity.Product{
		ID:    "second-product-id",
		Name:  "second-product-name",
		Price: 200,
	})

	productsList, err := listProductsUsecase.Execute()

	if err != nil {
		t.Error("Expected not to receive an error from ListProductsUsecase")
	}

	assertions.ExpectToBe(t, len(productsList), 2)
}
