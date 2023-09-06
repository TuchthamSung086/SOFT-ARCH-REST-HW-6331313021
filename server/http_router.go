package server

import (
	"softarchrest/handler"

	"github.com/gofiber/fiber/v2"
)

func (s *server) setRoutes(th *handler.TaskHandler) {
	// Define a route that responds with "Hello, World!" to all HTTP GET requests
	s.fiberApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	s.fiberApp.Post("/tasks", th.CreateTask)
	s.fiberApp.Get("/tasks", th.GetAllTasks)
}
