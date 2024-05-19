package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"net/http"
	"os"
	"short-urls/internal/config"
	"short-urls/internal/handlers/save"
	sqliteLog "short-urls/internal/lib/logger/sqlite"
	mwLogger "short-urls/internal/middleware/logger"
	"short-urls/internal/storage/sqlite"
)

func main() {
	cfg := config.MustLoad()

	log := initLogger(cfg.Env)
	log.Info("Starting app")
	log.Info("Debug level", slog.String("env", cfg.Env))

	storage, err := sqlite.New(cfg.DbStoragePath)
	if err != nil {
		log.Error("failed to init storage", sqliteLog.Err(err))
		os.Exit(1)
	}

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(mwLogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Post("/url", save.New(log, storage))

	log.Info("starting server", slog.String("address", cfg.Address))

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  4,
		WriteTimeout: 4,
		IdleTimeout:  30,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}

	log.Info("server stopped")
}

func initLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case config.EnvLocal, config.EnvDev:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case config.EnvProd:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
