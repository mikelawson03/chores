-- name: CreateChoreTemplate :exec
INSERT INTO chore_templates (id, name, cadence, assignee, duration, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: GetChoreTemplateByName :one
SELECT *
FROM chore_templates
WHERE name = ?;

-- name: GetChoreTemplateByID :one
SELECT *
FROM chore_templates
WHERE id = ?;

-- name: GetAllChoreTemplates :many
SELECT *
FROM chore_templates;

-- name: EditChoreTemplate :exec
UPDATE chore_templates
SET name = ?,
cadence = ?,
assignee = ?,
duration = ?,
updated_at = ?
WHERE id = ?;

-- name: DeleteChoreTemplate :exec
DELETE FROM chore_templates
WHERE id = ?;