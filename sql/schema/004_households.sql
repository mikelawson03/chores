-- +goose Up

CREATE TABLE households (
    id TEXT PRIMARY KEY,
    name TEXT UNIQUE
);

CREATE TABLE household_users (
    household_id TEXT NOT NULL,
    user_id TEXT NOT NULL,
    role TEXT NOT NULL,
    display_name TEXT,
    color_option INTEGER NOT NULL,
    joined_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,

    PRIMARY KEY(household_id, user_id),
    FOREIGN KEY (household_id) REFERENCES households(id),
    FOREIGN KEY (user_id) REFERENCES users(id),

    UNIQUE (household_id, color_option),
    CHECK (color_option >= 0 AND color_option < 10)
);

INSERT INTO households (id, name) 
VALUES (
    "e340f0e2-d10c-434c-90b4-a3b4c9028fed",
    "My House"
);

INSERT INTO household_users (
    household_id, 
    user_id, 
    role, 
    color_option
)
SELECT
    "e340f0e2-d10c-434c-90b4-a3b4c9028fed",
    id,
    role,
    ROW_NUMBER() OVER (
        ORDER BY created_at, id
    ) - 1
FROM users;

ALTER TABLE users DROP COLUMN role;

-- +goose Down
ALTER TABLE users ADD COLUMN role TEXT;

UPDATE users
SET role = (
    SELECT household_users.role
    FROM household_users
    WHERE household_users.user_id = users.id
);

DROP TABLE household_users;
DROP TABLE households;

