package auth

import (
	"github.com/go-kit/log"
	"github.com/jmoiron/sqlx"
)

type AuthService interface{}

type authService struct {
	logger log.Logger
}

func NewAuthService(logger log.Logger, db *sqlx.DB) AuthService {
	return &authService{
		logger: logger,
	}
}
