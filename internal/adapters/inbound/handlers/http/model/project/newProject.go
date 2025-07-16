package project

import ucModel "github.com/raihandotmd/taskify/internal/usecase/models/project"

type NewProjectRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

func (r *NewProjectRequest) ToUcModel() ucModel.NewProjectRequest {
	return ucModel.NewProjectRequest{
		Name:        r.Name,
		Description: r.Description,
	}
}
