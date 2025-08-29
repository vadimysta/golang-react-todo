package main

import (
	"context"
	"log"
	"time"

	"github.com/vadimysta/golang-react-todo/backend/internal/api"
	"github.com/vadimysta/golang-react-todo/backend/internal/config"
	"github.com/vadimysta/golang-react-todo/backend/internal/logger"
	"github.com/vadimysta/golang-react-todo/backend/internal/storage/mongodb"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	cfg, err := config.ConfigMustLoad()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	log := logger.SetupLogger(cfg)

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(10*time.Second))
	defer cancel()
	
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Error("invalid connects to URL the MongoDB", "error", err)
		return
	}

	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Error("failed to disconnect to MongoDB", "error", err)
			return
		}
	}()

	if err := client.Ping(ctx, nil); err != nil {
		log.Error("failed to connections MongoDB", "error", err)
		return
	}

	log.Info("Successfully to connections MongoDB")

	repo := mongodb.NewRepository(client)

	log.Info("Successfully to creat task")

	log.Info("Config successfully launched", "config", cfg)

	app := api.New(cfg, log, repo)

	log.Info("Application initialized")

	if err := app.Start(); err != nil {
		log.Error("Error starts server", "error", err)
	}

	log.Info("Server started successfully")

}