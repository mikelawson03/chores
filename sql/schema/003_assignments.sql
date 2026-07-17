-- +goose Up
CREATE TABLE assignments (
    id TEXT PRIMARY KEY,
    template_id TEXT NOT NULL,
    assigned_user_id TEXT NOT NULL,
    due_date TIMESTAMP NOT NULL,
    scheduled_for TIMESTAMP,
    completed BOOLEAN NOT NULL DEFAULT FALSE,
    canceled BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    completed_at TIMESTAMP,
    canceled_at TIMESTAMP,
    FOREIGN KEY (template_id) REFERENCES chore_templates(id),
    FOREIGN KEY (assigned_user_id) REFERENCES users(id)
);

-- +goose Down
DROP TABLE assignments;