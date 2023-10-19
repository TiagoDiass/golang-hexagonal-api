package entity

import (
	"testing"

	"github.com/TiagoDiass/golang-hexagonal-api/internal/test/assertions"
)

func TestNewProduct(t *testing.T) {
	product := NewProduct("iPhone 14", 2500)

	assertions.ExpectToBe(t, product.Name, "iPhone 14")
	assertions.ExpectToBe(t, product.Price, 2500)
}
