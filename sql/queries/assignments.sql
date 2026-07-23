-- name: CreateAssignment :exec
INSERT INTO assignments (
    id, 
    template_id, 
    assigned_user_id,
    due_date, 
    scheduled_for,
    instructions,
    notes,
    completed, 
    canceled, 
    created_at, 
    updated_at,
    completed_at,
    canceled_at)
    VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetAllAssignments :many
SELECT 
    a.*, 
    ct.name,
    ct.duration, 
    ct.cadence
FROM assignments a
JOIN chore_templates ct ON a.template_id = ct.id;

-- name: GetAssignment :one
SELECT 
    a.*, 
    ct.name,
    ct.duration, 
    ct.cadence
FROM assignments a
JOIN chore_templates ct ON a.template_id = ct.id
WHERE a.id = ?;

-- name: EditAssignment :exec
UPDATE assignments
SET assigned_user_id = ?,
scheduled_for = ?,
notes = ?,
created_at = ?,
updated_at = ?,
completed = ?,
canceled = ?,
completed_at = ?,
canceled_at = ?
WHERE id = ?;

-- name: GetAssignmentsByDateRange :many
SELECT *
FROM assignments
WHERE due_date >= ?
AND due_date < ?;

-- name: GetAssignmentsByTemplateID :many
SELECT *
FROM assignments
WHERE template_id = ?;

-- name: GetAssignmentsByUserID :many
SELECT *
FROM assignments
WHERE assigned_user_id = ?;

-- name: GetAssignmentsWithMetadataForDateRange :many
SELECT a.*, ct.duration, ct.cadence
FROM assignments a
JOIN chore_templates ct ON a.template_id = ct.id
WHERE (ct.cadence IN ("daily", "weekly") 
    AND a.due_date >= ?
    AND a.due_date < ?
    AND a.canceled = false)
OR (ct.cadence = "monthly"
    AND a.due_date >= ?
    AND a.due_date < ?
    AND a.canceled = false);