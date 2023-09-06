package handler

import (
	"fmt"
	"softarchrest/service"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Task struct {
	Title       string
	Description string
	DueDate     time.Time
	IsCompleted bool
}

func (t *Task) toServiceTask() *service.Task {
	return &service.Task{
		Title:       t.Title,
		Description: t.Description,
		DueDate:     t.DueDate,
		IsCompleted: t.IsCompleted,
	}
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
		return fmt.Errorf("TaskHandler CreateTask failed: %w", err)
	}

	// Return a response indicating success
	return ctx.Status(fiber.StatusCreated).JSON(newTask)
}

func (th *TaskHandler) GetAllTasks(ctx *fiber.Ctx) error {
	tasks, err := th.ts.GetAllTasks(ctx.Context())
	if err != nil {
		return fmt.Errorf("TaskHandler GetAllTasks failed: %w", err)
	}
	// Return the list of tasks as JSON response
	return ctx.JSON(tasks)
}

func (th *TaskHandler) GetTaskByID(ctx *fiber.Ctx) error {
	// Get the task ID from the URL parameter
	taskIDParam := ctx.Params("id")

	// Convert the task ID to an integer
	taskID, err := strconv.Atoi(taskIDParam)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid task ID",
		})
	}

	// Find the task with the specified ID
	task, err := th.ts.GetTaskByID(ctx.Context(), int32(taskID))
	if err != nil {
		return fmt.Errorf("TaskHandler GetTaskByID failed: %w", err)
	}

	// Check if the task was found
	if task == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Task not found",
		})
	}

	// Return the found task as JSON response
	return ctx.JSON(task)
}

func (th *TaskHandler) UpdateTask(ctx *fiber.Ctx) error {
	// Get the task ID from the URL parameter
	taskIDParam := ctx.Params("id")

	// Convert the task ID to an integer
	taskID, err := strconv.Atoi(taskIDParam)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid task ID",
		})
	}

	// Parse the JSON request body into a Task struct
	var newTask Task
	if err := ctx.BodyParser(&newTask); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Flag the fields that needs changing
	var x map[string]interface{}
	if err := ctx.BodyParser(&x); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	flags, err := th.getTaskUpdateFlags(x)
	if err != nil {
		return fmt.Errorf("TaskHandler getTaskUpdateFlags failed: %w", err)
	}

	// Update the task with the specified ID
	task, err := th.ts.UpdateTask(ctx.Context(), int32(taskID), *newTask.toServiceTask(), flags)
	if err != nil {
		return fmt.Errorf("TaskHandler UpdateTask failed: %w", err)
	}

	// Return the found task as JSON response
	return ctx.JSON(task)
}

func (th *TaskHandler) DeleteTask(ctx *fiber.Ctx) error {
	// Get the task ID from the URL parameter
	taskIDParam := ctx.Params("id")

	// Convert the task ID to an integer
	taskID, err := strconv.Atoi(taskIDParam)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid task ID",
		})
	}

	deletedTask, err := th.ts.DeleteTask(ctx.Context(), int32(taskID))
	if err != nil {
		return fmt.Errorf("TaskHandler DeleteTask failed: %w", err)
	}
	return ctx.JSON(deletedTask)
}

// TODO: this shit is cursed, but im too lazy to fix
func (th *TaskHandler) getTaskUpdateFlags(x map[string]interface{}) (*service.TaskUpdateFlags, error) {
	var flags service.TaskUpdateFlags

	for key, _ := range x {
		switch key {
		case "Title":
			flags.Title = true
		case "Description":
			flags.Description = true
		case "DueDate":
			flags.DueDate = true
		case "IsCompleted":
			flags.IsCompleted = true
		}
	}
	return &flags, nil
}
