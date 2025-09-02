package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/vadimysta/golang-react-todo/backend/internal/model"
	"github.com/vadimysta/golang-react-todo/backend/internal/storage/mongodb"
)

type Handler struct {
	todos mongodb.TodoRepository
}

func New(todos mongodb.TodoRepository) *Handler {
	return &Handler{
		todos: todos,
	}
}

func (h *Handler) CreatTodo(c *fiber.Ctx) error {

	var todo model.Todo

	if err := c.BodyParser(&todo); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid JSON"})
	}

	if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Todo body cannot be empty"})
	}

	if err := h.todos.Create(c.Context(), &todo); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to creat task"})
	}

	todo.Completed = false

	return c.Status(201).JSON(todo)

}

func (h *Handler) UpdateTodo(c *fiber.Ctx) error {

	id := c.Params("id")
	if id == "" {
		c.Status(400).JSON(fiber.Map{"error": "id not found"})
	}

	var request struct {
		Completed bool `json:"completed"`
	}

	if err := c.BodyParser(&request); err != nil {
		c.Status(400).JSON(fiber.Map{"error": "invalid JSON"})
	}

	if err := h.todos.Update(c.Context(), id, request.Completed); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to update task"})
	}
	

	return c.Status(200).JSON(fiber.Map{"message": "task update"})

}

func (h *Handler) DeleteTodo(c *fiber.Ctx) error {

	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"error": "id not found"})
	}

	if err := h.todos.Delete(c.Context(), id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to delete task"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "task delete"})
	
}

func (h *Handler) GetAllTodo(c *fiber.Ctx) error {
	
	todos, err := h.todos.GetAll(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to get task"})
	}

	return c.Status(201).JSON(todos)

}