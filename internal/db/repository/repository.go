package repository

import (
	"fmt"
	"go-fiber-api-docker/internal/db/models"

	"github.com/jmoiron/sqlx"
)

// type Repository interface {
// 	GetProductID(id int) (*models.Product, error)
// 	GetAllProduct() ([]*models.Product, error)
// 	UpdateProduct(id int, updated *models.Product) error
// }

type SQLRepository struct {
	db *sqlx.DB
}

func NewSQLRepository(db *sqlx.DB) *SQLRepository {
	return &SQLRepository{db: db}
}

func (r *SQLRepository) GetProductID(id int) (*models.Product, error) {
	query := "SELECT id, name, description, price, stock FROM products WHERE id = $1"
	row := r.db.QueryRow(query, id)

	fmt.Println(id)

	var product models.Product
	if err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock); err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *SQLRepository) GetAllProduct() ([]*models.Product, error) {
	query := "SELECT id, name, description, price, stock FROM products"
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	var products []*models.Product

	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return products, nil
}

func (r *SQLRepository) UpdateProduct(id int, updated *models.Product) error {
	query := "UPDATE products SET name = $1, description = $2, price = $3, stock = $4 WHERE id = $5"
	_, err := r.db.Exec(query, updated.Name, updated.Description, updated.Price, updated.Stock, id)
	return err
}

func (r *SQLRepository) DeleteProduct(id int) error {
	query := "DELETE FROM products WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return err
}

func (r *SQLRepository) AddProduct(product *models.Product) error {
	query := "INSERT INTO products (name, description, price, stock) VALUES ($1, $2, $3, $4)"
	_, err := r.db.Exec(query, product.Name, product.Description, product.Price, product.Stock)
	return err
}
