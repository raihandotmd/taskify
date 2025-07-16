package task

import ucModel "github.com/raihandotmd/taskify/internal/usecase/models/task"

type EditTaskRequest struct {
	ID          string `param:"id" validate:"required"`
	Title       string `json:"title" validate:"min=5"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Deadline    string `json:"deadline"`
}

func (r *EditTaskRequest) ToUcModel() ucModel.EditTaskRequest {
	return ucModel.EditTaskRequest{
		ID:          r.ID,
		Title:       r.Title,
		Description: r.Description,
		Status:      r.Status,
		Deadline:    r.Deadline,
	}
}
