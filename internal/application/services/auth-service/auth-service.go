package auth_service

import (
	"context"
	"github.com/nzb3/cakes-go/internal/application/models"
	"github.com/nzb3/cakes-go/internal/lib/logger"
)

type repository interface {
	GetUser(ctx context.Context, login string) (*models.User, error)
	GetUsers(ctx context.Context) ([]*models.User, error)
}

type service struct {
	repository repository
	log        logger.Logger
}

func NewService(log logger.Logger, repository repository) *service {
	return &service{
		repository: repository,
		log:        log,
	}
}

func (s *service) Login(ctx context.Context, login, password string) (*models.User, error) {
	// TODO: implement
	return nil, nil
}

func (s *service) Refresh(ctx context.Context, refreshToken string) (*models.User, error) {
	// TODO: implement
	return nil, nil
}