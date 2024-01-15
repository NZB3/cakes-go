package auth_controller

import (
	"context"
	"github.com/nzb3/cakes-go/internal/application/models"
	"github.com/nzb3/cakes-go/internal/lib/logger"
)

type authService interface {
	Login(ctx context.Context, login, password string) (*models.User, error)
	Refresh(ctx context.Context, refreshToken string) (*models.User, error)
}

type controller struct {
	log      logger.Logger
	services authService
}

func NewController(log logger.Logger, services authService) *controller {
	return &controller{
		log:      log,
		services: services,
	}
}
