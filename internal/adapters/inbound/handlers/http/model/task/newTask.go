package task

import ucModel "github.com/raihandotmd/taskify/internal/usecase/models/task"

type NewTaskRequest struct {
	ProjectID   string `json:"project_id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Status      string `json:"status" validate:"oneof=todo in_progress done,required"`
	Deadline    string `json:"deadline" validate:"required,datetime=2006-01-02"`
}

func (r *NewTaskRequest) ToUcModel() ucModel.NewTaskRequest {
	return ucModel.NewTaskRequest{
		ProjectID:   r.ProjectID,
		Title:       r.Title,
		Description: r.Description,
		Status:      r.Status,
		Deadline:    r.Deadline,
	}
}
