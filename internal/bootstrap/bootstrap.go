package bootstrap

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ash3in/go-cloudnative-boilerplate/internal/infra/config"
	"github.com/ash3in/go-cloudnative-boilerplate/internal/infra/logger"
	"github.com/ash3in/go-cloudnative-boilerplate/internal/infra/router"
)

type App struct {
	Logger *slog.Logger
	Server *http.Server
}

// Initialize loads configurations, initializes logger, sets up the router and HTTP server
func Initialize(ctx context.Context) (*App, error) {
	cfg := config.Load()

	log := logger.New()
	log.Info("logger initialized", "env", cfg.Env)

	mux := router.New(ctx)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.HTTPPort),
		Handler: mux,
	}

	app := &App{
		Logger: log,
		Server: srv,
	}

	return app, nil
}
