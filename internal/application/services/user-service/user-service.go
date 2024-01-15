package user_service

import (
	"context"
	"github.com/nzb3/cakes-go/internal/application/models"
	"github.com/nzb3/cakes-go/internal/lib/logger"
)

type repository interface {
	GetUser(ctx context.Context, login string) (*models.User, error)
	GetAllUsers(ctx context.Context) ([]models.User, error)
	CreateUser(ctx context.Context, user *models.User) error
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, login string) error
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

func (s *service) GetUser(ctx context.Context, login string) (*models.User, error) {
	errChan := make(chan error)
	userChan := make(chan *models.User)

	go func() {
		select {
		case <-ctx.Done():
			errChan <- ctx.Err()
		default:
			user, err := s.repository.GetUser(ctx, login)
			if err != nil {
				errChan <- err
				return
			}

			userChan <- user
		}
	}()

	select {
	case err := <-errChan:
		return nil, err
	case user := <-userChan:
		return user, nil
	}
}

func (s *service) GetUsers(ctx context.Context) ([]models.User, error) {
	errChan := make(chan error)
	usersChan := make(chan []models.User)

	go func() {
		select {
		case <-ctx.Done():
			errChan <- ctx.Err()
		default:
			users, err := s.repository.GetAllUsers(ctx)
			if err != nil {
				errChan <- err
				return
			}

			usersChan <- users
		}
	}()

	select {
	case err := <-errChan:
		return nil, err
	case users := <-usersChan:
		return users, nil
	}
}

func (s *service) CreateUser(ctx context.Context, user *models.User) error {
	errChan := make(chan error)

	go func() {
		select {
		case <-ctx.Done():
			errChan <- ctx.Err()
		default:
			errChan <- s.repository.CreateUser(ctx, user)
		}
	}()

	return <-errChan
}

func (s *service) UpdateUser(ctx context.Context, user *models.User) error {
	errChan := make(chan error)

	go func() {
		select {
		case <-ctx.Done():
			errChan <- ctx.Err()
		default:
			errChan <- s.repository.UpdateUser(ctx, user)
		}
	}()

	return <-errChan
}

func (s *service) DeleteUser(ctx context.Context, login string) error {
	errChan := make(chan error)

	go func() {
		select {
		case <-ctx.Done():
			errChan <- ctx.Err()
		default:
			errChan <- s.repository.DeleteUser(ctx, login)
		}
	}()

	return <-errChan
}
