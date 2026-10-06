-- name: AddUserToHousehold :exec
INSERT INTO household_users (household_id, user_id, role, display_name, color_option, joined_at, is_active)
VALUES (?, ?, ?, ?, ?, ?, ?);

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

-- name: GetHouseholdsForUser :many
SELECT household_id
FROM household_users
WHERE user_id = ?;

-- name: SetHouseholdUserActive :exec
UPDATE household_users
SET is_active = true
WHERE household_id = ?
AND user_id = ?;

-- name: SetHouseholdUserInactive :exec
UPDATE household_users
SET is_active = false
WHERE household_id = ?
AND user_id = ?;

-- name: EditHouseholdUser :one
UPDATE household_users
SET role = ?,
display_name = ?,
color_option = ?,
is_active = ?
WHERE user_id = ?
AND household_id = ?
RETURNING *;

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

-- name: DeleteHouseholdUser :execresult
DELETE FROM household_users
WHERE household_id = ?
AND user_id = ?;