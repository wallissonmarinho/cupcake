package cart

import (
	"github.com/go-kit/log"

	"github.com/jmoiron/sqlx"
)

type CartService interface{}

type cartService struct {
	logger log.Logger
}

func NewCartService(logger log.Logger, db *sqlx.DB) CartService {
	return &cartService{
		logger: logger,
	}
}
