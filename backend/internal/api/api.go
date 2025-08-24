package api

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/vadimysta/golang-react-todo/backend/internal/config"
	handlers "github.com/vadimysta/golang-react-todo/backend/internal/handler"
)

type App struct {
	Fiber  *fiber.App
	Config *config.Config
	Logger *slog.Logger
	Handler *handlers.Handler
}

func New(config *config.Config, log *slog.Logger) *App {

	app := App{
		Fiber:  fiber.New(),
		Config: config,
		Logger: log,
		Handler: handlers.New(),
	}

	app.SetupMiddleware()

	app.SetupRouters()


	return &app

}

func (a *App) SetupMiddleware() {

	a.Fiber.Use(fiberLogger.New())

}

func (a *App) SetupRouters() {

	a.Fiber.Get("/", a.Handler.HelloHandler)

}

func (a *App) Start() error {

	addr := a.Config.HTTPServer.Address

	a.Logger.Info("Starts server", "address", addr)

	return a.Fiber.Listen(addr)

}