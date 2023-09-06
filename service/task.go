package service

import (
	"context"
	"fmt"
	"softarchrest/database/gen"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Task struct {
	Title       string
	Description string
	DueDate     time.Time
	IsCompleted bool
}

type TaskUpdateFlags struct {
	Title       bool
	Description bool
	DueDate     bool
	IsCompleted bool
}

type TaskService struct {
	q *gen.Queries
}

func NewTaskService(q *gen.Queries) *TaskService {
	return &TaskService{q: q}
}

func (t *Task) fromGenTask(genTask *gen.Task) *Task {
	t.Title = genTask.Title
	t.Description = genTask.Description.String
	t.DueDate = genTask.DueDate.Time
	t.IsCompleted = genTask.IsCompleted.Bool
	return t
}

func (t *Task) toGenCreateTaskParams() *gen.CreateTaskParams {
	return &gen.CreateTaskParams{
		Title:       t.Title,
		Description: pgtype.Text{String: t.Description, Valid: true},
		DueDate:     pgtype.Date{Time: t.DueDate, Valid: true},
		IsCompleted: pgtype.Bool{Bool: t.IsCompleted, Valid: true},
	}
}

func (t *Task) toGenUpdateTaskParams(id int32) *gen.UpdateTaskParams {
	return &gen.UpdateTaskParams{
		ID:          id,
		Title:       t.Title,
		Description: pgtype.Text{String: t.Description, Valid: true},
		DueDate:     pgtype.Date{Time: t.DueDate, Valid: true},
		IsCompleted: pgtype.Bool{Bool: t.IsCompleted, Valid: true},
	}
}

func (ts *TaskService) CreateTask(ctx context.Context, newTask Task) error {
	err := ts.q.CreateTask(ctx, *newTask.toGenCreateTaskParams())
	if err != nil {
		return fmt.Errorf("TaskService CreateTask failed: %w", err)
	}
	return nil
}

func (ts *TaskService) GetAllTasks(ctx context.Context) ([]*Task, error) {
	genTasks, err := ts.q.GetAllTasks(ctx)
	if err != nil {
		return nil, fmt.Errorf("TaskService GetAllTasks failed: %w", err)
	}
	var tasks []*Task
	for _, genTask := range genTasks {
		var task Task
		tasks = append(tasks, task.fromGenTask(&genTask))
	}
	return tasks, nil
}

func (ts *TaskService) GetTaskByID(ctx context.Context, id int32) (*Task, error) {
	genTask, err := ts.q.GetTaskByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("TaskService GetTaskByID failed: %w", err)
	}
	var task Task
	return task.fromGenTask(&genTask), nil
}

func (ts *TaskService) UpdateTask(ctx context.Context, id int32, updatedTask Task, flags *TaskUpdateFlags) (*Task, error) {
	// Get original task
	currentTask, err := ts.GetTaskByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("TaskService UpdateTask failed: %w", err)
	}

	// For all that does not need changing, set to the original values
	if !flags.Title {
		updatedTask.Title = currentTask.Title
	}
	if !flags.Description {
		updatedTask.Description = currentTask.Description
	}
	if !flags.DueDate {
		updatedTask.DueDate = currentTask.DueDate
	}
	if !flags.IsCompleted {
		updatedTask.IsCompleted = currentTask.IsCompleted
	}

	err = ts.q.UpdateTask(ctx, *updatedTask.toGenUpdateTaskParams(id))
	if err != nil {
		return nil, fmt.Errorf("TaskService UpdateTask failed: %w", err)
	}
	return ts.GetTaskByID(ctx, id)
}

func (ts *TaskService) DeleteTask(ctx context.Context, id int32) (*Task, error) {
	targetTask, err := ts.GetTaskByID(ctx, id)

	err = ts.q.DeleteTask(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("TaskService DeleteTask failed: %w", err)
	}

	return targetTask, nil

}
