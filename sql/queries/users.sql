-- name: CreateUser :one
INSERT INTO users (id, username, password_hash, first_name, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: GetUserByUsername :one
SELECT id,
    username,
    first_name,
    created_at,
    updated_at
FROM users
WHERE username = ?;

-- name: GetUserByID :one
SELECT id,
    username,
    first_name,
    created_at,
    updated_at
FROM users
WHERE id = ?;

-- name: EditUser :one
UPDATE users
SET username = ?,
first_name =?,
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

-- name: GetHashForID :one
SELECT  *
FROM users
WHERE id = ?;

-- name: UpdatePassword :exec
UPDATE users
SET 
    password_hash = ?,
    updated_at = ?
WHERE id = ?;

-- name: UserHasHouseholdMemberships :one
SELECT EXISTS (
    SELECT 1
    FROM household_users
    WHERE user_id = ?
);

-- name: UserHasAssignments :one
SELECT EXISTS (
    SELECT 1
    FROM assignments
    WHERE assigned_user_id = ?
);