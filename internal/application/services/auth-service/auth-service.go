package auth_service

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nzb3/cakes-go/internal/application/models"
	auth_errors "github.com/nzb3/cakes-go/internal/errors/auth"
	"github.com/nzb3/cakes-go/internal/lib/logger"
	"net/http"
	"strings"
)

var JWT_SECRET = []byte("1234")

type repository interface {
	GetUser(ctx context.Context, login string) (*models.User, error)
	GetAllUsers(ctx context.Context) ([]models.User, error)
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
	user, err := s.repository.GetUser(ctx, login)
	if err != nil {
		return nil, nil
	}

	if user.Password != password {
		return nil, nil
	}
	return user, nil
}

func (s *service) Refresh(ctx context.Context, refreshToken string) (*models.User, error) {
	// TODO: implement
	return nil, nil
}

func (s *service) GetUserInfo(ctx context.Context, login string) (*models.User, error) {
	user, err := s.repository.GetUser(ctx, login)
	if err != nil {
		return nil, nil
	}
	return user, nil
}

func (s *service) AuthenticateUserWithToken(w http.ResponseWriter, r *http.Request) (*models.User, error) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return nil, auth_errors.New(w, "No Authorization Header", http.StatusBadRequest)
	}
	bearer := strings.Split(tokenString, " ")[1]

	token, err := jwt.Parse(bearer, func(token *jwt.Token) (interface{}, error) {
		return JWT_SECRET, nil
	})

	switch {
	case token.Valid:
		sub, _ := token.Claims.GetSubject()
		user, _ := s.GetUserInfo(r.Context(), sub)
		if user == nil {
			return nil, auth_errors.New(w, "User not founded", http.StatusNotFound)
		}
		return user, nil

	case errors.Is(err, jwt.ErrTokenExpired):
		return nil, auth_errors.New(w, "Token has expired", http.StatusUnauthorized)
	case errors.Is(err, jwt.ErrSignatureInvalid):
		return nil, auth_errors.New(w, "Signature is invalid", http.StatusBadRequest)
	}
	return nil, nil
}
