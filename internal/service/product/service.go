package product

import (
	"github.com/go-kit/log"
	"github.com/jmoiron/sqlx"
)

type ProductService interface{}

type productService struct {
	logger log.Logger
}

func NewProductService(logger log.Logger, db *sqlx.DB) ProductService {
	return &productService{
		logger: logger,
	}
}
