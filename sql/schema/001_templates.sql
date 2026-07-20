-- +goose Up

CREATE TABLE chore_templates (
    id TEXT PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    cadence TEXT NOT NULL,
    assignee TEXT,
    instructions TEXT,
    duration INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    FOREIGN KEY (assignee) REFERENCES users(id)
);

-- +goose Down
DROP TABLE chore_templates;