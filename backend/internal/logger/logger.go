package logger

import (
	"log/slog"
	"os"

	"github.com/vadimysta/golang-react-todo/backend/internal/config"
)


func SetupLogger(cfg *config.Config) *slog.Logger {

	var handler slog.Handler
	var logger *slog.Logger

	handlerOptsLocal := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	handlerOptsDev := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	handlerOptsProd := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	switch cfg.Env {
		case "local":
			handler = slog.NewTextHandler(os.Stdout, handlerOptsLocal)
		case "dev":
			handler = slog.NewJSONHandler(os.Stdout, handlerOptsDev)
		case "prod":
			handler = slog.NewJSONHandler(os.Stdout, handlerOptsProd)
		default:
			handler = slog.NewTextHandler(os.Stdout, handlerOptsLocal)
	}

	logger = slog.New(handler)

	return logger

}