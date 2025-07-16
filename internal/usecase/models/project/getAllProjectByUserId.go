package project

import (
	obModel "github.com/raihandotmd/taskify/internal/adapters/outbound/models/project"
)

type (
	GetAllProjectByUserIdResponse struct {
		Projects []obModel.Project `json:"projects"`
	}
)
