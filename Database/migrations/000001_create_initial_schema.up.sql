CREATE TYPE task_type AS ENUM (
    'boolean',
    'number',
    'duration',
    'time'
    );

-- Create Users table
CREATE TABLE IF NOT EXISTS users
(
    user_uuid  UUID PRIMARY KEY         DEFAULT gen_random_uuid(),
    email      VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    archived   BOOLEAN                  DEFAULT FALSE
);

-- Create tasks table
CREATE TABLE IF NOT EXISTS tasks
(
    task_uuid   UUID PRIMARY KEY         DEFAULT gen_random_uuid(),
    user_uuid   UUID         NOT NULL REFERENCES users (user_uuid) ON DELETE CASCADE,
    type        task_type    NOT NULL,
    name        VARCHAR(255) NOT NULL,
    description TEXT,
    archived    BOOLEAN                  DEFAULT FALSE,
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS progress
(
    progress_uuid  UUID PRIMARY KEY,
    task_uuid      UUID      NOT NULL REFERENCES tasks (task_uuid) ON DELETE CASCADE,
    progress_type  task_type NOT NULL,
    value          JSONB     NOT NULL,
    created_at_utc TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at_utc TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create index for faster queries by user
CREATE INDEX IF NOT EXISTS idx_tasks_user ON tasks (user_uuid);
-- Create a GIN index for efficient JSONB querying
CREATE INDEX IF NOT EXISTS idx_progress_value ON progress USING GIN (value);
-- Create additional indexes for common query patterns
CREATE INDEX IF NOT EXISTS idx_tasks_archived ON tasks (archived);
CREATE INDEX IF NOT EXISTS idx_tasks_type ON tasks (type);
CREATE INDEX IF NOT EXISTS idx_progress_task_uuid ON progress (task_uuid);
