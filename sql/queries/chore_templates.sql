-- name: CreateChoreTemplate :exec
INSERT INTO chore_templates (id, name, cadence, assignee, instructions, duration, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?);

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
instructions =?,
duration = ?,
updated_at = ?
WHERE id = ?;

-- name: DeleteChoreTemplate :one
DELETE FROM chore_templates
WHERE id = ?
RETURNING id;