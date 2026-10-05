package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/iDIOtu/PivchikiDND/internal/config"
	"github.com/iDIOtu/PivchikiDND/internal/database"
	"github.com/iDIOtu/PivchikiDND/internal/delivery"
)

func main() {
	logger := slog.New(
		slog.NewTextHandler(
			os.Stdout,
			nil,
		),
	)

	slog.SetDefault(logger)

	cfg := config.Load()

	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	defer stop()

	db, err := database.NewPostgres(
		ctx,
		cfg.DatabaseURL,
	)
	if err != nil {
		logger.Error(
			"failed to connect to postgres",
			"error",
			err,
		)

		os.Exit(1)
	}

	defer db.Close()

	logger.Info(
		"postgres connection established",
	)

	handler := delivery.NewHandler(db)

	router := delivery.AppRouter(
		handler,
	)

	server := &http.Server{
		Addr:    cfg.AppAddr,
		Handler: router,

		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	go func() {
		logger.Info(
			"server started",
			"addr",
			cfg.AppAddr,
		)

		if err := server.ListenAndServe(); err != nil &&
			!errors.Is(
				err,
				http.ErrServerClosed,
			) {
			logger.Error(
				"server failed",
				"error",
				err,
			)

			stop()
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)

	defer cancel()

	if err := server.Shutdown(
		shutdownCtx,
	); err != nil {
		logger.Error(
			"server shutdown failed",
			"error",
			err,
		)
	}
}