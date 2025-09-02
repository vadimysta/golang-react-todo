package api

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/vadimysta/golang-react-todo/backend/internal/config"
	handlers "github.com/vadimysta/golang-react-todo/backend/internal/handler"
	"github.com/vadimysta/golang-react-todo/backend/internal/storage/mongodb"
)

type App struct {
	Fiber  *fiber.App
	Config *config.Config
	Logger *slog.Logger
	Handler *handlers.Handler
}

func New(config *config.Config, log *slog.Logger, repo *mongodb.TodoRepository) *App {

	// repo := 

	app := App{
		Fiber:  fiber.New(),
		Config: config,
		Logger: log,
		Handler: handlers.New(*repo),
	}

	app.SetupCORS()

	app.SetupMiddleware()

	app.SetupRouters()


	return &app

}

func (a *App) SetupMiddleware() {

	a.Fiber.Use(fiberLogger.New())

}

func (a *App) SetupCORS() {

	a.Fiber.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowMethods: "GET, POST, PATCH, DELETE",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

}

func (a *App) SetupRouters() {

	a.Fiber.Post("/todo", a.Handler.CreatTodo)
	a.Fiber.Patch("/todo/:id", a.Handler.UpdateTodo)
	a.Fiber.Delete("/todo/:id", a.Handler.DeleteTodo)
	a.Fiber.Get("/todo", a.Handler.GetAllTodo)

}

func (a *App) Start() error {

	addr := a.Config.HTTPServer.Address

	a.Logger.Info("Starts server", "address", addr)

	return a.Fiber.Listen(addr)

}