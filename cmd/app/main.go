package main

import (
	"database/sql"
	"net/http"

	repositoryimpl "github.com/TiagoDiass/golang-hexagonal-api/internal/infra/repository_impl"
	"github.com/TiagoDiass/golang-hexagonal-api/internal/infra/web"
	"github.com/TiagoDiass/golang-hexagonal-api/internal/usecase"
	"github.com/go-chi/chi/v5"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306)/products")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	productRepository := repositoryimpl.NewProductRepositoryMysql(db)
	createProductUsecase := usecase.NewCreateProductUsecase(productRepository)
	listProductsUsecase := usecase.NewListProductsUsecase(productRepository)
	productHandlers := web.NewProductHandlers(createProductUsecase, listProductsUsecase)

	router := chi.NewRouter()

	router.Get("/products", productHandlers.ListProductsHandler)
	router.Post("/products", productHandlers.CreateProductHandler)

	http.ListenAndServe(":8000", router)
}
