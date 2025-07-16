package project

import ucModel "github.com/raihandotmd/taskify/internal/usecase/models/project"

type EditProjectRequest struct {
	ID          string `param:"id" validate:"required"`
	Name        string `json:"name" validate:"min=5"`
	Description string `json:"description"`
}

func (r *EditProjectRequest) ToUcModel() ucModel.EditProjectRequest {
	return ucModel.EditProjectRequest{
		ID:          r.ID,
		Name:        r.Name,
		Description: r.Description,
	}
}
