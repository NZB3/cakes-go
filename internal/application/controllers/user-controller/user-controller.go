package user_controller

import (
	"context"
	"encoding/json"
	"github.com/nzb3/cakes-go/internal/application/models"
	"github.com/nzb3/cakes-go/internal/lib/logger"
	"net/http"
	"os"
	"path"
)

type userService interface {
	GetUser(ctx context.Context, login string) (*models.User, error)
	GetUsers(ctx context.Context) ([]models.User, error)
	CreateUser(ctx context.Context, user *models.User) error
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, login string) error
}

type controller struct {
	log      logger.Logger
	services userService
}

func NewController(log logger.Logger, services userService) *controller {
	return &controller{
		log:      log,
		services: services,
	}
}

func (c *controller) UserHandler(w http.ResponseWriter, r *http.Request) {
	c.log.Infof("handle at: %s", r.URL.Path)

	prefix := os.Getenv("API_PREFIX")

	switch path.Join(prefix, path.Base(r.URL.Path)) {
	case prefix + "/users":
		c.getAllUsersHandler(w, r)
	case prefix + "/user":
		c.handleMethod(w, r)
	default:
		http.Error(w, "Not found", http.StatusNotFound)
	}
}

func (c *controller) handleMethod(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		c.getUserHandler(w, r)
	case http.MethodPost:
		c.addUserHandler(w, r)
	case http.MethodPut:
		c.updateUserHandler(w, r)
	case http.MethodDelete:
		c.deleteUserHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (c *controller) getAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	c.log.Infof("handle get all users at: %s", r.URL.Path)

	users, err := c.services.GetUsers(r.Context())
	if err != nil {
		c.log.Errorf("error handling get users: %w", err)

		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if len(users) == 0 {
		http.Error(w, "Not users found", http.StatusNotFound)
		return
	}

	js, err := json.Marshal(users)
	if err != nil {
		c.log.Errorf("error handling marshal users: %w", err)

		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(js); err != nil {
		c.log.Errorf("error handling write response: %w", err)
	}
}

func (c *controller) getUserHandler(w http.ResponseWriter, r *http.Request) {
	c.log.Infof("handle get user at: %s", r.URL.Path)

	login := r.URL.Query().Get("login")
	if login == "" {
		http.Error(w, "Login is required", http.StatusBadRequest)
		return
	}

	user, err := c.services.GetUser(r.Context(), login)
	if err != nil {
		c.log.Errorf("error handling get user: %w", err)

		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(user)
	if err != nil {
		c.log.Errorf("error handling marshal user: %w", err)

		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(js); err != nil {
		c.log.Errorf("error handling write response: %w", err)
	}
}

func (c *controller) addUserHandler(w http.ResponseWriter, r *http.Request) {
	c.log.Infof("handle add user at: %s", r.URL.Path)

	user := models.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		c.log.Errorf("error handling decode user: %w", err)

		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if err := c.services.CreateUser(r.Context(), &user); err != nil {
		c.log.Errorf("error handling create user: %w", err)

		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *controller) updateUserHandler(w http.ResponseWriter, r *http.Request) {
	c.log.Infof("handle update user at: %s", r.URL.Path)

	user := models.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		c.log.Errorf("error handling decode user: %w", err)

		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if err := c.services.UpdateUser(r.Context(), &user); err != nil {
		c.log.Errorf("error handling update user: %w", err)

		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *controller) deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	c.log.Infof("handle delete user at: %s", r.URL.Path)

	var login string
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		c.log.Errorf("error handling decode user: %w", err)

		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if err := c.services.DeleteUser(r.Context(), login); err != nil {
		c.log.Errorf("error handling delete user: %w", err)

		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
