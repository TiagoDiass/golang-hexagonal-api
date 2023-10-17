package web

import (
	"encoding/json"
	"net/http"

	"github.com/TiagoDiass/golang-hexagonal-api/internal/usecase"
	"github.com/go-chi/chi/v5"
)

type ProductHandlers struct {
	CreateProductUsecase   *usecase.CreateProductUsecase
	ListProductsUsecase    *usecase.ListProductsUsecase
	DeleteProductUsecase   *usecase.DeleteProductUsecase
	FindProductByIdUsecase *usecase.FindProductByIdUsecase
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewProductHandlers(
	createProductUsecase *usecase.CreateProductUsecase,
	listProductsUsecase *usecase.ListProductsUsecase,
	deleteProductUsecase *usecase.DeleteProductUsecase,
	findProductByIdUsecase *usecase.FindProductByIdUsecase,
) *ProductHandlers {
	return &ProductHandlers{
		CreateProductUsecase:   createProductUsecase,
		ListProductsUsecase:    listProductsUsecase,
		DeleteProductUsecase:   deleteProductUsecase,
		FindProductByIdUsecase: findProductByIdUsecase,
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

func (h *ProductHandlers) DeleteProductHandler(response http.ResponseWriter, request *http.Request) {
	productId := chi.URLParam(request, "productId")

	var input usecase.DeleteProductInputDTO = usecase.DeleteProductInputDTO{
		ID: productId,
	}

	err := h.DeleteProductUsecase.Execute(input)

	if err != nil {
		errorResponse := ErrorResponse{
			Error: err.Error(),
		}

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(errorResponse)

		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusNoContent)
}

func (h *ProductHandlers) FindProductByIdHandler(response http.ResponseWriter, request *http.Request) {
	productId := chi.URLParam(request, "productId")

	input := usecase.FindProductByIdInputDTO{
		ID: productId,
	}

	product, err := h.FindProductByIdUsecase.Execute(input)

	if err != nil {
		errorResponse := ErrorResponse{
			Error: err.Error(),
		}

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(errorResponse)

		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(product)
}
