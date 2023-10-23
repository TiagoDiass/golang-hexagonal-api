package main

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"

	repositoryimpl "github.com/TiagoDiass/golang-hexagonal-api/internal/infra/repository_impl"
	"github.com/TiagoDiass/golang-hexagonal-api/internal/infra/web"
	"github.com/TiagoDiass/golang-hexagonal-api/internal/usecase"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306)/products")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	CreateProductsTableIfNotExists(db)

	productHandlers := CreateProductHandlers(db)
	router := InitializeRoutes(productHandlers)

	http.ListenAndServe(":8000", router)
}

func CreateProductsTableIfNotExists(db *sql.DB) {
	db.Exec(`create table if not exists products (
		id varchar(255), 
		name varchar(255), 
		price integer,
		PRIMARY KEY (id)
	);`)
}

func CreateProductHandlers(db *sql.DB) *web.ProductHandlers {
	productRepository := repositoryimpl.NewProductRepositoryMysql(db)
	createProductUsecase := usecase.NewCreateProductUsecase(productRepository)
	listProductsUsecase := usecase.NewListProductsUsecase(productRepository)
	deleteProductUsecase := usecase.NewDeleteProductUsecase(productRepository)
	findProductByIdUsecase := usecase.NewFindProductByIdUsecase(productRepository)

	productHandlers := web.NewProductHandlers(
		createProductUsecase,
		listProductsUsecase,
		deleteProductUsecase,
		findProductByIdUsecase,
	)

	return productHandlers
}

func InitializeRoutes(productHandlers *web.ProductHandlers) *chi.Mux {
	router := chi.NewRouter()

	router.Get("/products", productHandlers.ListProductsHandler)
	router.Post("/products", productHandlers.CreateProductHandler)
	router.Get("/products/{productId}", productHandlers.FindProductByIdHandler)
	router.Delete("/products/{productId}", productHandlers.DeleteProductHandler)

	return router
}
