package dashboard

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type DashboardService interface{}

type dashboardService struct {
	logger *log.Logger
}

func NewDashboardService(logger *log.Logger, db *sqlx.DB) DashboardService {
	return &dashboardService{
		logger: logger,
	}
}
