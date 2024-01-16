package repository

import (
	"context"
	"github.com/nzb3/cakes-go/internal/application/models"
)

func (r *repository) GetIngredient(ctx context.Context, article int) (*models.Ingredient, error) {
	var ingredient models.Ingredient

	err := r.db.WithContext(ctx).Where("article = ?", article).First(&ingredient).Error
	if err != nil {
		r.log.Errorf("error getting ingredient: %w", err)
		return nil, err
	}

	return &ingredient, nil
}

func (r *repository) GetAllIngredients(ctx context.Context) ([]models.Ingredient, error) {
	ingredients := make([]models.Ingredient, 0)

	err := r.db.WithContext(ctx).Find(&ingredients).Error
	if err != nil {
		r.log.Errorf("error getting ingredients: %w", err)
		return nil, err
	}

	return ingredients, nil
}

func (r *repository) CreateIngredient(ctx context.Context, ingredient *models.Ingredient) error {
	err := r.db.WithContext(ctx).Create(ingredient).Error
	if err != nil {
		r.log.Errorf("error creating ingredient: %w", err)
		return err
	}

	return nil
}

func (r *repository) UpdateIngredient(ctx context.Context, ingredient *models.Ingredient) error {
	err := r.db.WithContext(ctx).Save(ingredient).Error
	if err != nil {
		r.log.Errorf("error updating ingredient: %w", err)
		return err
	}

	return nil
}

func (r *repository) DeleteIngredient(ctx context.Context, article int) error {
	err := r.db.WithContext(ctx).Where("article = ?", article).Delete(&models.Ingredient{}).Error
	if err != nil {
		r.log.Errorf("error deleting ingredient: %w", err)
		return err
	}

	return nil
}
