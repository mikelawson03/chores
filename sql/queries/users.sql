-- name: CreateUser :one

INSERT INTO users (id, username, role, first_name, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: GetUserByUsername :one
SELECT *
FROM users
WHERE username = ?;

-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = ?;

-- name: GetAllUsers :many
SELECT *
FROM users;

-- name: EditUser :one
UPDATE users
SET username = ?,
first_name =?,
role = ?,
updated_at = ?
WHERE id = ?
RETURNING *;

-- name: DeleteUser :one
DELETE FROM users
WHERE id = ?
RETURNING id;