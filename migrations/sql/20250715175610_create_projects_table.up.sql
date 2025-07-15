CREATE TABLE projects (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(255) NOT NULL,
  description TEXT,
  created_by UUID NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  
  CONSTRAINT fk_projects_created_by 
    FOREIGN KEY (created_by) 
    REFERENCES users(id) 
    ON DELETE CASCADE
);

-- Create index on created_by for faster lookups
CREATE INDEX idx_projects_created_by ON projects(created_by);

-- Create index on created_at for sorting
CREATE INDEX idx_projects_created_at ON projects(created_at);