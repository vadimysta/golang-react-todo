package main

import (
	"log"

	"github.com/vadimysta/golang-react-todo/backend/internal/api"
	"github.com/vadimysta/golang-react-todo/backend/internal/config"
	"github.com/vadimysta/golang-react-todo/backend/internal/logger"
)

func main() {

	cfg, err := config.ConfigMustLoad()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	log := logger.SetupLogger(cfg)

	log.Info("Config successfully launched", "config", cfg)

	app := api.New(cfg, log)

	log.Info("Application initialized")

	if err := app.Start(); err != nil {
		log.Error("Error starts server", "error", err)
	}

	log.Info("Server started successfully")

}