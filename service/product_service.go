package service

import (
	"context"
	"database/sql"
	"restful-api/dto/request"
	"restful-api/dto/response"
	"restful-api/helper"
	"restful-api/repository"

	"github.com/go-playground/validator/v10"
)

type ProductService struct {
	DB                *sql.DB
	productRepository repository.ProductRepository
	validate          *validator.Validate
}

func NewProductService(db *sql.DB, productRepository repository.ProductRepository, validate *validator.Validate) *ProductService {
	return &ProductService{
		DB:                db,
		productRepository: productRepository,
		validate:          validate,
	}
}

func (ps *ProductService) CreateProduct(ctx context.Context, request request.CreateProductRequest) response.CreateProductResponse {
	err := ps.validate.Struct(request)

	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			return response.CreateProductResponse{
				Status: 403,
				Errors: helper.ExtractValidationErrors(validationErrors, request),
			}
		}
	}

	tx, err := ps.DB.BeginTx(ctx, nil)

	if err != nil {
		return response.CreateProductResponse{
			Status:  500,
			Message: "Internal Server Error",
		}
	}
	defer tx.Rollback()

	result, err := tx.ExecContext(ctx, "INSERT INTO products (name, type) VALUES (?,?)", request.Name, request.Type)

	if err != nil {
		return response.CreateProductResponse{
			Status:  500,
			Message: "Internal Server Error",
		}
	}

	lastInsertedId, err := result.LastInsertId()

	if err != nil {
		return response.CreateProductResponse{
			Status:  500,
			Message: "Internal Server Error",
		}
	}

	err = tx.Commit()

	if err != nil {
		return response.CreateProductResponse{
			Status:  500,
			Message: "Internal Server Error",
		}
	}

	return response.CreateProductResponse{
		Status:  201,
		Message: "Created",
		Data: response.ProductData{
			ProductId:   int(lastInsertedId),
			ProductName: request.Name,
			ProductType: request.Name,
		},
	}
}
