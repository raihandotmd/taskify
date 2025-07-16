package project

import (
	obModel "github.com/raihandotmd/taskify/internal/adapters/outbound/models/project"
)

type (
	NewProjectRequest struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	NewProjectResponse struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		UserID      string `json:"user_id"`
	}
)

func (req NewProjectRequest) ToObModel(userId string) obModel.Project {
	return obModel.Project{
		Name:        req.Name,
		Description: req.Description,
		UserID:      userId,
	}
}
