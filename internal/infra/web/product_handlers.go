package web

import (
	"encoding/json"
	"net/http"

	"github.com/TiagoDiass/golang-hexagonal-api/internal/usecase"
)

type ProductHandlers struct {
	CreateProductUsecase *usecase.CreateProductUsecase
	ListProductsUsecase  *usecase.ListProductsUsecase
}

func NewProductHandlers(createProductUsecase *usecase.CreateProductUsecase, listProductsUsecase *usecase.ListProductsUsecase) *ProductHandlers {
	return &ProductHandlers{
		CreateProductUsecase: createProductUsecase,
		ListProductsUsecase:  listProductsUsecase,
	}
}

func (h *ProductHandlers) CreateProductHandler(response http.ResponseWriter, request *http.Request) {
	var input usecase.CreateProductInputDTO

	err := json.NewDecoder(request.Body).Decode(&input)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := h.CreateProductUsecase.Execute(input)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(output)
}

func (h *ProductHandlers) ListProductsHandler(response http.ResponseWriter, request *http.Request) {
	output, err := h.ListProductsUsecase.Execute()

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(output)
}
