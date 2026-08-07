-- name: CreateUser :one

INSERT INTO users (id, username, password_hash, role, first_name, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: GetUserByUsername :one
SELECT id,
    username,
    role,
    first_name,
    created_at,
    updated_at
FROM users
WHERE username = ?;

-- name: GetUserByID :one
SELECT id,
    username,
    role,
    first_name,
    created_at,
    updated_at
FROM users
WHERE id = ?;

-- name: GetAllUsers :many
SELECT id,
    username,
    role,
    first_name,
    created_at,
    updated_at
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

-- name: GetUserCount :one
SELECT COUNT(*) FROM users;

-- name: GetHashForUsername :one
SELECT *
FROM users
WHERE username = ?;

-- name: UpdatePassword :exec
UPDATE users
SET password_hash = ?
WHERE id = ?