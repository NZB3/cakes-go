package auth_controller

import (
	"context"
	"encoding/json"
	"github.com/nzb3/cakes-go/internal/application/models"
	"github.com/nzb3/cakes-go/internal/lib/logger"
	"net/http"
	"os"
)

type authService interface {
	Login(ctx context.Context, login, password string) (*models.User, error)
	Refresh(ctx context.Context, refreshToken string) (*models.User, error)
}

type controller struct {
	log      logger.Logger
	services authService
}

func NewController(log logger.Logger, services authService) *controller {
	return &controller{
		log:      log,
		services: services,
	}
}

func (c *controller) AuthHandler(w http.ResponseWriter, r *http.Request) {
	//	todo switch case structure that implementing all urls
	prefix := os.Getenv("API_PREFIX")

	switch r.URL.Path {
	case prefix + "/auth":
		switch r.Method {
		case http.MethodPost:
			c.LoginUser(w, r)
		default:
			http.Error(w, "Method Not allowed", http.StatusMethodNotAllowed)
		}

	}
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *controller) LoginUser(w http.ResponseWriter, r *http.Request) {
	var p []byte
	var request LoginRequest

	r.Body.Read(p)
	err := json.Unmarshal(p, &request)

	if err == nil {
		http.Error(w, "Bad Request. Unparsed json", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\": \"ok\"}"))
}
