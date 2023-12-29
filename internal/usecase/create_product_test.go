package usecase

import (
	"testing"

	"github.com/TiagoDiass/golang-hexagonal-api/internal/test"
	"github.com/TiagoDiass/golang-hexagonal-api/internal/test/assertions"
)

func TestShouldCreateProductSuccessfully(t *testing.T) {
	t.Parallel()

	productRepository := test.NewInMemoryProductRepository()
	createProductUsecase := NewCreateProductUsecase(productRepository)

	input := CreateProductInputDTO{
		Name:  "fake-product-name",
		Price: 1500,
	}

	output, err := createProductUsecase.Execute(input)

	if err != nil {
		t.Error("Expected not to receive an error from CreateProductUsecase")
	}

	assertions.ExpectToBe(t, output.Name, "fake-product-name")
	assertions.ExpectToBe(t, output.Price, 1500)
}
