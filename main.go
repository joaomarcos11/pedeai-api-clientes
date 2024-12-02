package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"

	"github.com/filipeandrade6/fiap-pedeai-clientes/adapters/repository/postgresql"
	"github.com/filipeandrade6/fiap-pedeai-clientes/controllers/api"
	"github.com/filipeandrade6/fiap-pedeai-clientes/domain/services"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// ====================
	// database

	db, err := postgresql.New(ctx, postgresql.Config{
		Host:       os.Getenv("DB_HOST"),
		Port:       os.Getenv("DB_PORT"),
		User:       os.Getenv("DB_USER"),
		Password:   os.Getenv("DB_PASS"),
		Name:       os.Getenv("DB_NAME"),
		DisableTLS: false,
	})
	if err != nil {
		logger.Error("connecting to database", "error", err)
	}

	srv := api.NewServer(logger, services.New(db))

	httpServer := &http.Server{
		Addr:    ":8081",
		Handler: srv,
	}

	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Info("listening and serving", "error", err)
	}
}
