-- name: CreateUser :exec

INSERT INTO users (id, username, role, created_at, updated_at)
VALUES (?, ?, ?, ?, ?);

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

-- name: EditUser :exec
UPDATE users
SET username = ?,
updated_at = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;