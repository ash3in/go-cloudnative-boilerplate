package main

import (
	"context"
	"os/signal"
	"syscall"
	"time"

	"github.com/ash3in/go-cloudnative-boilerplate/internal/bootstrap"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	app, err := bootstrap.Initialize(ctx)
	if err != nil {
		panic("failed to bootstrap: " + err.Error())
	}

	go func() {
		app.Logger.Info("starting server", "addr", app.Server.Addr)
		if err := app.Server.ListenAndServe(); err != nil && err.Error() != "http: server closed" {
			app.Logger.Error("server failed", "error", err)
			stop()
		}
	}()

	<-ctx.Done()
	app.Logger.Info("shutdown initiated")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.Server.Shutdown(shutdownCtx); err != nil {
		app.Logger.Error("graceful shutdown failed", "error", err)
	} else {
		app.Logger.Info("server stopped gracefully")
	}
}
