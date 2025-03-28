-- Drop indexes
DROP INDEX IF EXISTS idx_progress_task_uuid;
DROP INDEX IF EXISTS idx_tasks_type;
DROP INDEX IF EXISTS idx_tasks_archived;
DROP INDEX IF EXISTS idx_progress_value;
DROP INDEX IF EXISTS idx_tasks_user;

-- Drop tables in reverse order (to avoid foreign key constraints)
DROP TABLE IF EXISTS progress;
DROP TABLE IF EXISTS tasks;
DROP TABLE IF EXISTS users;

-- Drop types
DROP TYPE IF EXISTS task_type;
