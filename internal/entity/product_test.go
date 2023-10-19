package entity

import "testing"

func ExpectToBe(t *testing.T, value interface{}, expectedValue interface{}) {
	errorMessage := "\nExpected: %v\nReceived: %v"

	if value != expectedValue {
		t.Errorf(errorMessage, expectedValue, value)
	}
}

func TestNewProduct(t *testing.T) {
	product := NewProduct("iPhone 14", 2500)

	ExpectToBe(t, product.Name, "iPhone 14")
	ExpectToBe(t, product.Price, 2500)
}
