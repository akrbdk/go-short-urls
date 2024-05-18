package main

import (
	"log/slog"
	"os"
	"short-urls/internal/config"
	sqliteLog "short-urls/internal/logger/sqlite"
	"short-urls/internal/storage/sqlite"
)

func main() {
	//TODO init config
	cfg := config.MustLoad()

	//TODO init logger
	log := initLogger(cfg.Env)
	log.Info("Starting app")
	log.Info("Debug level", slog.String("env", cfg.Env))

	//TODO init storage
	storage, err := sqlite.New(cfg.DbStoragePath)
	if err != nil {
		log.Error("failed to init storage", sqliteLog.Err(err))
		os.Exit(1)
	}

	_ = storage

	//TODO init router
	//TODO run app
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
