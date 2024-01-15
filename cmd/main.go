package main

import (
	"context"
	"github.com/nzb3/cakes-go/internal/application"
	"github.com/nzb3/cakes-go/internal/config"
	"github.com/nzb3/cakes-go/internal/lib/logger"
	"os"
	"os/signal"
	"syscall"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	ctx := context.Background()

	config.MustLoadConfig(".env")

	log := logger.NewLogger()

	app := application.NewApp(log)

	go func() {
		app.Run(ctx)
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	app.Stop(ctx)

	log.Info("Gracefully stopped")

}
