-- name: CreateTemplate :exec
INSERT INTO chore_templates (id, name, cadence, shared, assignee, duration, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?);