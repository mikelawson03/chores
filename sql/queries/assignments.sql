-- name: CreateAssignment :exec
INSERT INTO assignments (
    id, 
    template_id, 
    assigned_user_id, 
    scheduled_for, 
    completed, 
    canceled, 
    created_at, 
    updated_at,
    completed_at,
    canceled_at)
    VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetAllAssignments :many
SELECT *
FROM assignments;

-- name: GetAssignmentByID :one
SELECT *
FROM assignments
WHERE id = ?;

-- name: EditAssignment :exec
UPDATE assignments
SET assigned_user_id = ?,
scheduled_for = ?,
created_at = ?,
updated_at = ?
WHERE id = ?;

-- name: CancelAssignment :exec
UPDATE assignments
SET canceled = ?,
canceled_at = ?,
updated_at = ?
WHERE id = ?;

-- name: CompleteAssignment :exec
UPDATE assignments
SET completed = ?,
completed_at = ?,
updated_at = ?
WHERE id = ?;

-- name: GetAssignmentsByDateRange :many
SELECT *
FROM assignments
WHERE scheduled_for >= ?
AND scheduled_for < ?;