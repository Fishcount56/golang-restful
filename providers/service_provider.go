package providers

import (
	"database/sql"
	"restful-api/service"

	"github.com/go-playground/validator/v10"
)

type ServiceProvider struct {
	ProductService *service.ProductService
}

func NewServiceProvider(db *sql.DB, validate *validator.Validate) *ServiceProvider {
	return &ServiceProvider{
		ProductService: service.NewProductService(db, validate),
	}
}
