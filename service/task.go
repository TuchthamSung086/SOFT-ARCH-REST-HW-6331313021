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

type TaskService struct {
	q *gen.Queries
}

func NewTaskService(q *gen.Queries) *TaskService {
	return &TaskService{q: q}
}

func (ts *TaskService) CreateTask(ctx context.Context, newTask Task) error {
	err := ts.q.CreateTask(ctx,
		gen.CreateTaskParams{
			Title:       newTask.Title,
			Description: pgtype.Text{String: newTask.Description, Valid: true},
			DueDate:     pgtype.Date{Time: newTask.DueDate, Valid: true},
			IsCompleted: pgtype.Bool{Bool: newTask.IsCompleted, Valid: true},
		})
	return err
}

func (ts *TaskService) GetAllTasks(ctx context.Context) ([]Task, error) {
	genTasks, err := ts.q.GetAllTasks(ctx)
	if err != nil {
		return nil, fmt.Errorf("Task Service failed: %w", err)
	}
	var tasks []Task
	for _, genTask := range genTasks {
		tasks = append(tasks, Task{
			Title:       genTask.Title,
			Description: genTask.Description.String,
			DueDate:     genTask.DueDate.Time,
			IsCompleted: genTask.IsCompleted.Bool,
		})
	}
	return tasks, nil
}
