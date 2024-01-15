package repository

import (
	"context"
	"github.com/nzb3/cakes-go/internal/application/models"
)

func (r *repository) GetAllCakeDecorations(ctx context.Context) ([]models.CakeDecoration, error) {
	var cakeDecorations []models.CakeDecoration
	err := r.db.WithContext(ctx).Find(&cakeDecorations).Error
	if err != nil {
		return nil, err
	}

	return cakeDecorations, nil
}

func (r *repository) GetCakeDecoration(ctx context.Context, article int) (*models.CakeDecoration, error) {
	var cakeDecoration models.CakeDecoration
	err := r.db.WithContext(ctx).Where("article = ?", article).First(&cakeDecoration).Error
	if err != nil {
		return nil, err
	}

	return &cakeDecoration, nil
}

func (r *repository) CreateCakeDecoration(ctx context.Context, cakeDecoration *models.CakeDecoration) error {
	err := r.db.WithContext(ctx).Create(cakeDecoration).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateCakeDecoration(ctx context.Context, cakeDecoration *models.CakeDecoration) error {
	err := r.db.WithContext(ctx).Save(cakeDecoration).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteCakeDecoration(ctx context.Context, article int) error {
	err := r.db.WithContext(ctx).Where("article = ?", article).Delete(&models.CakeDecoration{}).Error
	if err != nil {
		return err
	}

	return nil
}
