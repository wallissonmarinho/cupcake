package profile

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type ProfileService interface{}

type profileService struct {
	logger *log.Logger
}

func NewProfileService(logger *log.Logger, db *sqlx.DB) ProfileService {
	return &profileService{
		logger: logger,
	}
}
