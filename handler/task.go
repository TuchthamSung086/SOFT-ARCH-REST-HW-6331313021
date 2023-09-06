package handler

import (
	"fmt"
	"softarchrest/service"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Task struct {
	Title       string
	Description string
	DueDate     time.Time
	IsCompleted bool
}

type TaskHandler struct {
	ts *service.TaskService
}

func NewTaskHandler(ts *service.TaskService) *TaskHandler {
	return &TaskHandler{ts: ts}
}

func (th *TaskHandler) CreateTask(ctx *fiber.Ctx) error {
	// Parse the JSON request body into a Task struct
	var newTask Task
	if err := ctx.BodyParser(&newTask); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Perform any required logic to create the task (e.g., save it to a database)
	err := th.ts.CreateTask(ctx.Context(), service.Task{
		Title:       newTask.Title,
		Description: newTask.Description,
		DueDate:     newTask.DueDate,
		IsCompleted: newTask.IsCompleted,
	})
	if err != nil {
		return fmt.Errorf("Task Handler failed: %w", err)
	}

	// Return a response indicating success
	return ctx.Status(fiber.StatusCreated).JSON(newTask)
}

func (th *TaskHandler) GetAllTasks(ctx *fiber.Ctx) error {
	tasks, err := th.ts.GetAllTasks(ctx.Context())
	if err != nil {
		return fmt.Errorf("Task Handler failed: %w", err)
	}
	// Return the list of tasks as JSON response
	return ctx.JSON(tasks)
}
