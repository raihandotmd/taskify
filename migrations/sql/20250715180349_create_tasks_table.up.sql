CREATE TYPE task_status AS ENUM ('todo', 'in_progress', 'done');

CREATE TABLE tasks (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  project_id UUID NOT NULL,
  title VARCHAR(255) NOT NULL,
  description TEXT,
  status task_status NOT NULL DEFAULT 'todo',
  deadline DATE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  
  CONSTRAINT fk_tasks_project_id 
    FOREIGN KEY (project_id) 
    REFERENCES projects(id) 
    ON DELETE CASCADE
);

-- Create index on project_id for faster lookups
CREATE INDEX idx_tasks_project_id ON tasks(project_id);

-- Create index on status for filtering
CREATE INDEX idx_tasks_status ON tasks(status);

-- Create index on deadline for sorting
CREATE INDEX idx_tasks_deadline ON tasks(deadline);

-- Create index on created_at for sorting
CREATE INDEX idx_tasks_created_at ON tasks(created_at);