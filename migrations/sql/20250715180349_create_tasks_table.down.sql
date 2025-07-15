DROP INDEX IF EXISTS idx_tasks_created_at;
DROP INDEX IF EXISTS idx_tasks_deadline;
DROP INDEX IF EXISTS idx_tasks_status;
DROP INDEX IF EXISTS idx_tasks_project_id;
DROP TABLE IF EXISTS tasks;
DROP TYPE IF EXISTS task_status;