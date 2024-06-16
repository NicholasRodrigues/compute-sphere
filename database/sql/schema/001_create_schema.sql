-- +goose Up
CREATE TABLE users
(
    id            UUID PRIMARY KEY,
    username      VARCHAR(255) UNIQUE      NOT NULL,
    password_hash VARCHAR(255)             NOT NULL,
    role          VARCHAR(50)              NOT NULL CHECK (role IN ('admin', 'volunteer')),
    created_at    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE tasks
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(255)             NOT NULL,
    payload     JSONB,
    parameters  JSONB,
    status      VARCHAR(50)              NOT NULL CHECK (status IN ('pending', 'in-progress', 'completed', 'failed')),
    assigned_to INT REFERENCES users (id),
    priority    INT                               DEFAULT 0,
    retries     INT                               DEFAULT 0,
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE results
(
    id         SERIAL PRIMARY KEY,
    task_id    INT REFERENCES tasks (id) ON DELETE CASCADE,
    user_id    INT REFERENCES users (id),
    output     JSONB,
    elapsed_time INTERVAL,
    status     VARCHAR(50)              NOT NULL CHECK (status IN ('success', 'failure')),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_tasks_status ON tasks (status);
CREATE INDEX idx_results_task_id ON results (task_id);
CREATE INDEX idx_results_user_id ON results (user_id);


-- +goose Down
DROP TABLE results;
DROP TABLE tasks;
DROP TABLE users;
