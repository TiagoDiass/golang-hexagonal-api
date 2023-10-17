package repositoryimpl

import (
	"database/sql"

	"github.com/TiagoDiass/golang-hexagonal-api/internal/entity"
)

type ProductRepositoryMysql struct {
	DB *sql.DB
}

func NewProductRepositoryMysql(db *sql.DB) *ProductRepositoryMysql {
	return &ProductRepositoryMysql{DB: db}
}

func (r *ProductRepositoryMysql) Create(product *entity.Product) error {
	_, err := r.DB.Exec(
		"insert into products (id, name, price) values(?,?,?)",
		product.ID, product.Name, product.Price,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *ProductRepositoryMysql) FindAll() ([]*entity.Product, error) {
	rows, err := r.DB.Query("select id, name, price from products")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product

	for rows.Next() {
		var product entity.Product

		err = rows.Scan(&product.ID, &product.Name, &product.Price)

		if err != nil {
			return nil, err
		}

		products = append(products, &product)
	}

	return products, nil
}

func (r *ProductRepositoryMysql) Delete(id string) error {
	_, err := r.DB.Exec("delete from products where id = ?", id)

	if err != nil {
		return err
	}

	return nil
}

func (r *ProductRepositoryMysql) FindById(id string) (*entity.Product, error) {
	rows, err := r.DB.Query("select id, name, price from products where id = ?", id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	var product entity.Product

	rows.Scan(&product.ID, &product.Name, &product.Price)

	return &product, nil
}
