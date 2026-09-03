-- name: CreateUser :one
INSERT INTO users (id, username, password_hash, first_name, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: AddUserToHousehold :exec
INSERT INTO household_users (household_id, user_id, role, display_name, color_option, joined_at, is_active)
VALUES (?, ?, ?, ?, ?, ?, ?);

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

-- name: GetHouseholdUsers :many
SELECT 
    hu.household_id,
    hu.role,
    hu.display_name,
    hu.color_option,
    hu.joined_at,
    hu.is_active,
    u.id,
    u.username,
    u.first_name,
    u.created_at,
    u.updated_at
FROM household_users hu
JOIN users u ON u.id = hu.user_id
WHERE hu.household_id = ?
ORDER BY hu.joined_at, u.id;

-- name: GetHouseholdUserByID :one
SELECT
    hu.household_id,
    hu.role,
    hu.display_name,
    hu.color_option,
    hu.joined_at,
    hu.is_active,
    u.id,
    u.username,
    u.first_name,
    u.created_at,
    u.updated_at
FROM household_users hu
JOIN users u ON u.id = hu.user_id
WHERE hu.household_id = ?
AND hu.user_id = ?;

-- name: EditUser :one
UPDATE users
SET username = ?,
first_name =?,
updated_at = ?
WHERE id = ?
RETURNING *;

-- name: EditHouseholdUser :one
UPDATE household_users
SET role = ?,
display_name = ?,
color_option = ?,
is_active = ?
WHERE user_id = ?
AND household_id = ?
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
WHERE id = ?;

-- name: HouseholdColorOptionInUse :one
SELECT EXISTS (
    SELECT 1
    FROM household_users
    WHERE household_id = ?
    AND color_option = ?
    AnD user_id != ?
);

-- name: HouseholdUsersCount :one
SELECT COUNT(*)
FROM household_users;
