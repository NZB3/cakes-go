package auth_controller

import (
	"context"
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nzb3/cakes-go/internal/application/models"
	"github.com/nzb3/cakes-go/internal/lib/logger"
	"net/http"
	"os"
	"time"
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

type JWTResponse struct {
	Token string `json:"token"`
}

func (c *controller) LoginUser(w http.ResponseWriter, r *http.Request) {
	var request LoginRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, "Bad Request. Unparsed json", http.StatusBadRequest)
		return
	}

	user, _ := c.services.Login(r.Context(), request.Username, request.Password)
	if user == nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	claims := jwt.MapClaims{
		"sub": user.Login,
		"exp": time.Now().Add(15 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenWithSecret, _ := token.SignedString([]byte("1234"))

	responseData, _ := json.Marshal(JWTResponse{Token: tokenWithSecret})

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseData)
}
