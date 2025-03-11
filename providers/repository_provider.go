package providers

import (
	"database/sql"
	"restful-api/repository"
)

type RepositoryProvider struct {
	ProductRepository *repository.ProductRepository
}

func NewRepositoryProvider(db *sql.DB) *RepositoryProvider {
	return &RepositoryProvider{
		ProductRepository: repository.NewProductRepository(db),
	}
}
