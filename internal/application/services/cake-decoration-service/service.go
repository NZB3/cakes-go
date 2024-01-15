package cake_decoration_service

import (
	"context"
	"github.com/nzb3/cakes-go/internal/application/models"
	"github.com/nzb3/cakes-go/internal/lib/logger"
)

type repository interface {
	GetAllCakeDecorations(ctx context.Context) ([]models.CakeDecoration, error)
	GetCakeDecoration(ctx context.Context, article int) (*models.CakeDecoration, error)
	CreateCakeDecoration(ctx context.Context, cakeDecoration *models.CakeDecoration) error
	UpdateCakeDecoration(ctx context.Context, cakeDecoration *models.CakeDecoration) error
	DeleteCakeDecoration(ctx context.Context, article int) error
}

type service struct {
	repository
	log logger.Logger
}

func NewService(log logger.Logger, repository repository) *service {
	return &service{
		repository: repository,
		log:        log,
	}
}

func (s *service) GetCakeDecorations(ctx context.Context) ([]models.CakeDecoration, error) {
	errChan := make(chan error, 1)
	defer close(errChan)

	cakeDecorationsChan := make(chan []models.CakeDecoration, 1)
	defer close(cakeDecorationsChan)

	go func() {
		select {
		case <-ctx.Done():
			errChan <- ctx.Err()
		default:
			cakeDecorations, err := s.repository.GetAllCakeDecorations(ctx)
			if err != nil {
				errChan <- err
				return
			}

			cakeDecorationsChan <- cakeDecorations
		}
	}()

	select {
	case err := <-errChan:
		return nil, err
	case cakeDecorations := <-cakeDecorationsChan:
		return cakeDecorations, nil
	}
}

func (s *service) GetCakeDecoration(ctx context.Context, article int) (*models.CakeDecoration, error) {
	errChan := make(chan error)
	cakeDecorationChan := make(chan *models.CakeDecoration)

	go func() {
		select {
		case <-ctx.Done():
			errChan <- ctx.Err()
		default:
			cakeDecoration, err := s.repository.GetCakeDecoration(ctx, article)
			if err != nil {
				errChan <- err
				return
			}

			cakeDecorationChan <- cakeDecoration
		}
	}()

	select {
	case err := <-errChan:
		return nil, err
	case cakeDecoration := <-cakeDecorationChan:
		return cakeDecoration, nil
	}
}

func (s *service) CreateCakeDecoration(ctx context.Context, cakeDecoration models.CakeDecoration) error {
	errChan := make(chan error, 1)
	defer close(errChan)

	go func() {
		select {
		case <-ctx.Done():
			errChan <- ctx.Err()
		default:
			errChan <- s.repository.CreateCakeDecoration(ctx, &cakeDecoration)
		}
	}()

	return <-errChan
}

func (s *service) UpdateCakeDecoration(ctx context.Context, cakeDecoration models.CakeDecoration) error {
	errChan := make(chan error, 1)
	defer close(errChan)

	go func() {
		select {
		case <-ctx.Done():
			errChan <- ctx.Err()
		default:
			errChan <- s.repository.UpdateCakeDecoration(ctx, &cakeDecoration)
		}
	}()

	return <-errChan
}

func (s *service) DeleteCakeDecoration(ctx context.Context, article int) error {
	errChan := make(chan error, 1)
	defer close(errChan)

	go func() {
		select {
		case <-ctx.Done():
			errChan <- ctx.Err()
		default:
			errChan <- s.repository.DeleteCakeDecoration(ctx, article)
		}
	}()

	return <-errChan
}
