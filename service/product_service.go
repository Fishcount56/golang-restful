package service

import (
	"database/sql"
	"restful-api/dto/request"
	"restful-api/dto/response"
	"restful-api/helper"

	"github.com/go-playground/validator/v10"
)

type ProductService struct {
	DB       *sql.DB
	validate *validator.Validate
}

func NewProductService(db *sql.DB, validate *validator.Validate) *ProductService {
	return &ProductService{
		DB:       db,
		validate: validate,
	}
}

func (ps *ProductService) CreateProduct(request request.CreateProductRequest) response.CreateProductResponse {
	err := ps.validate.Struct(request)

	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			return response.CreateProductResponse{
				Status: 403,
				Errors: helper.ExtractValidationErrors(validationErrors, request),
			}
		}
	}

	tx, err := ps.DB.Begin()

	if err != nil {
		return response.CreateProductResponse{
			Status:  500,
			Message: "Internal Server Error",
		}
	}
	defer tx.Rollback()

	result, err := tx.Exec("INSERT INTO products (name, type) VALUES (?,?)", request.Name, request.Type)

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
