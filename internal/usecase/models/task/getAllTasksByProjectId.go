package task

import (
	obModel "github.com/raihandotmd/taskify/internal/adapters/outbound/models/task"
)

type (
	GetAllTasksByProjectIdResponse struct {
		Tasks []obModel.Task `json:"tasks"`
	}
)
