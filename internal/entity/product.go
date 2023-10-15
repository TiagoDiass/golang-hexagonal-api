package entity

import "github.com/google/uuid"

type Product struct {
	ID    string
	Name  string
	Price int
}

func NewProduct(name string, price int) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}
