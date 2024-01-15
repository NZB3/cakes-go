package repository

import (
	"context"
	"github.com/nzb3/cakes-go/internal/application/models"
	"github.com/samber/lo"
)

type userDB struct {
	Login    string `db:"login"`
	Password string `db:"password"`
	Role     string `db:"role"`
	FullName string `db:"full_name"`
	Photo    string `db:"photo"`
}

func (s *repository) GetUser(ctx context.Context, login string) (*models.User, error) {
	var user userDB

	err := s.db.WithContext(ctx).Where("login = ?", login).First(&user).Error
	if err != nil {
		s.log.Errorf("error getting user: %w", err)
		return nil, err
	}

	return &models.User{
		Login:    user.Login,
		Password: user.Password,
		Role:     user.Role,
		FullName: user.FullName,
		Photo:    user.Photo,
	}, nil
}

func (s *repository) GetUsers(ctx context.Context) ([]*models.User, error) {
	users := make([]userDB, 0)

	err := s.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		s.log.Errorf("error getting users: %w", err)
		return nil, err
	}

	return lo.Map(users, func(user userDB, _ int) *models.User {
		return &models.User{
			Login:    user.Login,
			Password: user.Password,
			Role:     user.Role,
			FullName: user.FullName,
			Photo:    user.Photo,
		}
	}), nil
}

func (s *repository) CreateUser(ctx context.Context, user *models.User) error {
	err := s.db.WithContext(ctx).Create(user).Error
	if err != nil {
		s.log.Errorf("error creating user: %w", err)
		return err
	}

	return nil
}

func (s *repository) UpdateUser(ctx context.Context, user *models.User) error {
	err := s.db.WithContext(ctx).Save(user).Error
	if err != nil {
		s.log.Errorf("error updating user: %w", err)
		return err
	}

	return nil
}

func (s *repository) DeleteUser(ctx context.Context, login string) error {
	err := s.db.WithContext(ctx).Where("login = ?", login).Delete(&userDB{}).Error
	if err != nil {
		s.log.Errorf("error deleting user: %w", err)
		return err
	}

	return nil
}
