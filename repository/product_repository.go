package repository

import (
	"context"
	"database/sql"
	"restful-api/dto/request"
)

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

func (pr *ProductRepository) CreateProduct(ctx context.Context, tx *sql.Tx, request request.CreateProductRequest) (int64, error) {
	result, err := tx.ExecContext(ctx, "INSERT INTO products (name, type) VALUES (?, ?)", request.Name, request.Type)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}
