-- name: CreateTask :exec
INSERT INTO tasks (title, description, due_date, is_completed)
VALUES ($1, $2, $3, $4);
-- name: GetAllTasks :many
SELECT *
FROM tasks;
-- name: GetAllUnfinishedTasks :many
SELECT *
FROM tasks
WHERE is_completed = false;
-- name: GetTaskByID :one
SELECT *
FROM tasks
WHERE id = $1
LIMIT 1;
-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = $1;
-- name: UpdateTask :exec
UPDATE tasks
SET title = $2,
    description = $3,
    due_date = $4,
    is_completed = $5
WHERE id = $1;