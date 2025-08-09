package order

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type OrderService interface{}

type orderService struct {
	logger *log.Logger
}

func NewOrderService(logger *log.Logger, db *sqlx.DB) OrderService {
	return &orderService{
		logger: logger,
	}
}
