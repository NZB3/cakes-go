package repository

import (
	"context"
	"github.com/nzb3/cakes-go/internal/application/models"
)

func (r *repository) CreateTool(ctx context.Context, tool *models.Tool) error {
	err := r.db.WithContext(ctx).Create(tool).Error
	if err != nil {
		r.log.Errorf("error creating user: %w", err)
		return err
	}

	return nil
}

func (r *repository) GetTool(ctx context.Context, name string) (*models.Tool, error) {
	var tool models.Tool

	err := r.db.WithContext(ctx).Where("name = ?", name).First(&tool).Error
	if err != nil {
		r.log.Errorf("error getting user: %w", err)
		return nil, err
	}

	return &tool, nil
}

func (r *repository) GetTools(ctx context.Context) ([]*models.Tool, error) {
	tools := make([]*models.Tool, 0)

	err := r.db.WithContext(ctx).Find(&tools).Error
	if err != nil {
		r.log.Errorf("error getting users: %w", err)
		return nil, err
	}

	return tools, nil
}

func (r *repository) UpdateTool(ctx context.Context, tool *models.Tool) error {
	err := r.db.WithContext(ctx).Save(tool).Error
	if err != nil {
		r.log.Errorf("error updating user: %w", err)
		return err
	}

	return nil
}

func (r *repository) DeleteTool(ctx context.Context, name string) error {
	err := r.db.WithContext(ctx).Where("name = ?", name).Delete(&models.Tool{}).Error
	if err != nil {
		r.log.Errorf("error deleting user: %w", err)
		return err
	}

	return nil
}
