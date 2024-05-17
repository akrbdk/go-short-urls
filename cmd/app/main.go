package main

import (
	"log/slog"
	"os"
	"short-urls/internal/config"
)

func main() {
	//TODO init config
	cfg := config.MustLoad()

	//TODO init logger
	log := initLogger(cfg.Env)
	log.Info("Starting app")
	log.Info("Debug level", slog.String("env", cfg.Env))

	//TODO init storage
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
