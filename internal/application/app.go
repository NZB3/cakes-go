package application

import (
	"context"
	"fmt"
	main_controller "github.com/nzb3/cakes-go/internal/application/controllers/main-controller"
	user_controller "github.com/nzb3/cakes-go/internal/application/controllers/user-controller"
	"github.com/nzb3/cakes-go/internal/application/repository"
	user_service "github.com/nzb3/cakes-go/internal/application/services/user-service"
	"github.com/nzb3/cakes-go/internal/lib/logger"
	"github.com/nzb3/cakes-go/internal/lib/router"
	"net"
	"net/http"
	"os"
)

type app struct {
	server *http.Server
	log    logger.Logger
}

func NewApp(log logger.Logger) *app {
	return &app{
		log: log,
	}
}

func (a *app) Run(ctx context.Context) {
	repo := repository.NewRepository(a.log)

	userService := user_service.NewService(a.log, repo)
	//authService := auth_service.NewService(a.log, repo)

	mainController := main_controller.NewController(a.log)
	userController := user_controller.NewController(a.log, userService)
	// authController := auth_controller.NewController(a.log, authService)

	r := router.NewRouter()
	r.Mount("/", mainController.MainHandler)
	r.Mount("/users", userController.UserHandler)

	a.server = &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")),
		Handler: r.GetHandler(),
		BaseContext: func(net.Listener) context.Context {
			return ctx
		},
	}

	a.log.Info("Starting server")

	if err := a.server.ListenAndServe(); err != nil {
		a.log.Errorf("error starting server: %s", err.Error())
	}
}

func (a *app) Stop(ctx context.Context) {
	a.log.Info("Stopping server")
	if err := a.server.Shutdown(ctx); err != nil {
		a.log.Errorf("error stopping server: %s", err.Error())
	}
}
