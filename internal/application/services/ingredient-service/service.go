package ingredient_service

import (
	"context"
	"github.com/nzb3/cakes-go/internal/application/models"
)

type repository interface {
	GetAllIngredients(ctx context.Context) ([]models.Ingredient, error)
	GetIngredient(ctx context.Context, article int) (*models.Ingredient, error)
	CreateIngredient(ctx context.Context, ingredient *models.Ingredient) error
	UpdateIngredient(ctx context.Context, ingredient *models.Ingredient) error
	DeleteIngredient(ctx context.Context, article int) error
}

type service struct {
	repository
}

func NewService(repository repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetAllIngredients(ctx context.Context) ([]models.Ingredient, error) {
	errChan := make(chan error, 1)
	defer close(errChan)

	ingredientsChan := make(chan []models.Ingredient, 1)
	defer close(ingredientsChan)

	go func() {
		select {
		case <-ctx.Done():
			errChan <- ctx.Err()
		default:
			ingredients, err := s.repository.GetAllIngredients(ctx)
			if err != nil {
				errChan <- err
				return
			}

			ingredientsChan <- ingredients
		}
	}()

	select {
	case err := <-errChan:
		return nil, err
	case ingredients := <-ingredientsChan:
		return ingredients, nil
	}
}

func (s *service) GetIngredient(ctx context.Context, article int) (*models.Ingredient, error) {
	errChan := make(chan error)
	ingredientChan := make(chan *models.Ingredient)

	go func() {
		select {
		case <-ctx.Done():
			errChan <- ctx.Err()
		default:
			ingredient, err := s.repository.GetIngredient(ctx, article)
			if err != nil {
				errChan <- err
				return
			}

			ingredientChan <- ingredient
		}
	}()

	select {
	case err := <-errChan:
		return nil, err
	case ingredient := <-ingredientChan:
		return ingredient, nil
	}
}

func (s *service) CreateIngredient(ctx context.Context, ingredient *models.Ingredient) error {
	errChan := make(chan error)

	go func() {
		select {
		case <-ctx.Done():
			errChan <- ctx.Err()
		default:
			errChan <- s.repository.CreateIngredient(ctx, ingredient)
		}
	}()

	return <-errChan
}

func (s *service) UpdateIngredient(ctx context.Context, ingredient *models.Ingredient) error {
	errChan := make(chan error)

	go func() {
		select {
		case <-ctx.Done():
			errChan <- ctx.Err()
		default:
			errChan <- s.repository.UpdateIngredient(ctx, ingredient)
		}
	}()

	return <-errChan
}

func (s *service) DeleteIngredient(ctx context.Context, article int) error {
	errChan := make(chan error)

	go func() {
		select {
		case <-ctx.Done():
			errChan <- ctx.Err()
		default:
			errChan <- s.repository.DeleteIngredient(ctx, article)
		}
	}()

	return <-errChan
}
