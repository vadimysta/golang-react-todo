package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/vadimysta/golang-react-todo/backend/internal/model"
)

type Handler struct {
	todos []model.Todo
}

func New() *Handler {
	return &Handler{
		todos: []model.Todo{},
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

	todo.ID = len(h.todos) + 1 
	todo.Completed = false

	h.todos = append(h.todos, todo)

	return c.Status(201).JSON(todo)

}

func (h *Handler) UpdateTodo(c *fiber.Ctx) error {

	id := c.Params("id")

	for i, v := range h.todos {

		if fmt.Sprint(v.ID) == id {
			h.todos[i].Completed = true

			return c.Status(200).JSON(h.todos[i])
		
		}
	}

	return c.Status(404).JSON(fiber.Map{"error": "task not found"})

}

func (h *Handler) DeleteTodo(c *fiber.Ctx) error {

	id := c.Params("id")

	for i, v := range h.todos {

		if fmt.Sprint(v.ID) == id {
			h.todos = append(h.todos[:i], h.todos[i+1:]...)

			return c.Status(200).JSON(fiber.Map{"success": true})

		}
	} 

	return c.Status(404).JSON(fiber.Map{"error": "failed delete to task"})

}

func (h *Handler) GetAllTodo(c *fiber.Ctx) error {
	
	return c.JSON(h.todos)

}