package repository

import (
	"context"
	"github.com/nzb3/cakes-go/internal/application/models"
)

func (r *repository) GetUser(ctx context.Context, login string) (*models.User, error) {
	var user models.User

	err := r.db.WithContext(ctx).Where("login = ?", login).First(&user).Error
	if err != nil {
		r.log.Errorf("error getting user: %w", err)
		return nil, err
	}

	return &user, nil
}

func (r *repository) GetUsers(ctx context.Context) ([]models.User, error) {
	users := make([]models.User, 0)

	err := r.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		r.log.Errorf("error getting users: %w", err)
		return nil, err
	}

	return users, nil
}

func (r *repository) CreateUser(ctx context.Context, user *models.User) error {
	err := r.db.WithContext(ctx).Create(user).Error
	if err != nil {
		r.log.Errorf("error creating user: %w", err)
		return err
	}

	return nil
}

func (r *repository) UpdateUser(ctx context.Context, user *models.User) error {
	err := r.db.WithContext(ctx).Save(user).Error
	if err != nil {
		r.log.Errorf("error updating user: %w", err)
		return err
	}

	return nil
}

func (r *repository) DeleteUser(ctx context.Context, login string) error {
	err := r.db.WithContext(ctx).Where("login = ?", login).Delete(&models.User{}).Error
	if err != nil {
		r.log.Errorf("error deleting user: %w", err)
		return err
	}

	return nil
}
