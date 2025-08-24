package handlers

import "github.com/gofiber/fiber/v2"

type Handler struct {}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) HelloHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Hello serverssssss",
	})
}