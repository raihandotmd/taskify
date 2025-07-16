package project

import obModel "github.com/raihandotmd/taskify/internal/adapters/outbound/models/project"

type (
	EditProjectRequest struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	EditProjectResponse struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		UserID      string `json:"created_by"`
	}
)

func (req EditProjectRequest) ToObModel(userId string) obModel.Project {
	return obModel.Project{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description,
		UserID:      userId,
	}
}
