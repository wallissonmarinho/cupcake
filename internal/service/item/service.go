package item

import (
	"github.com/go-kit/log"

	"github.com/jmoiron/sqlx"
)

type ItemService interface{}

type itemService struct {
	logger log.Logger
}

func NewItemService(logger log.Logger, db *sqlx.DB) ItemService {
	return &itemService{
		logger: logger,
	}
}
