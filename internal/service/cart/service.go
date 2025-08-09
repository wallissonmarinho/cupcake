package cart

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type CartService interface{}

type cartService struct {
	logger *log.Logger
}

func NewCartService(logger *log.Logger, db *sqlx.DB) CartService {
	return &cartService{
		logger: logger,
	}
}
