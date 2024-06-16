-- name: AddTask :exec
INSERT INTO tasks (name, payload, parameters, status, assigned_to, priority, retries)
VALUES ($1, $2, $3, 'pending', $4, $5, $6);

-- name: GetPendingTask :one
SELECT id, name, payload, parameters, status, assigned_to, priority, retries, created_at, updated_at
FROM tasks
WHERE status = 'pending'
LIMIT 1;

-- name: UpdateTaskStatus :exec
UPDATE tasks
SET status = $2, updated_at = NOW()
WHERE id = $1;
