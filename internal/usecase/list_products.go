package usecase

import "github.com/TiagoDiass/golang-hexagonal-api/internal/repository"

type ListProductsOutputDTO struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type ListProductsUsecase struct {
	ProductRepository repository.ProductRepository
}

func NewListProductsUsecase(productRepository repository.ProductRepository) *ListProductsUsecase {
	return &ListProductsUsecase{
		ProductRepository: productRepository,
	}
}

func (u *ListProductsUsecase) Execute() ([]*ListProductsOutputDTO, error) {
	products, err := u.ProductRepository.FindAll()

	if err != nil {
		return nil, err
	}

	var productsOutput []*ListProductsOutputDTO

	for _, product := range products {
		productsOutput = append(productsOutput, &ListProductsOutputDTO{
			ID:    product.ID,
			Name:  product.Name,
			Price: product.Price,
		})
	}

	return productsOutput, nil
}
