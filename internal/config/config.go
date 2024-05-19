package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

const EnvLocal = "local"
const EnvDev = "dev"
const EnvProd = "prod"

type Config struct {
	Env           string `yaml:"env" env-default:"local"`
	DbStoragePath string `yaml:"db_storage_path" env-default:"./db"`
	HttpConfig
}

type HttpConfig struct {
	Address string `yaml:"address" env-default:"127.0.0.1:8080"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "./config/config.yaml"
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("CONFIG_PATH does not exist %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Error loading config: %s", err)
	}

	return &cfg
}
