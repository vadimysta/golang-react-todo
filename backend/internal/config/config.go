package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string     `yaml:"env" env:"ENV" default:"prod" env-required:"true"`
	HTTPServer HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
	Address string `yaml:"address" env:"ADDRESS" default:"localhost:8080" env-required:"true"`
	Host    string `yaml:"host" env:"HOST" default:"localhost"`
	Port    string `yaml:"port" env:"PORT" default:"8080"`
}

var cfg Config

const defaultEnvConfigPath = "CONFIG_PATH"

const op = "./backend/internal/config/config.go-ConfigMustLoad"

func ConfigMustLoad() (*Config, error) {

	configPath := os.Getenv(defaultEnvConfigPath)
	if configPath == "" {
		configPath = "./backend/config/config.yaml"
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("%s: %v", op, err)
	}

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, fmt.Errorf("%s: %v", op, err)
	}

	return &cfg, nil

}