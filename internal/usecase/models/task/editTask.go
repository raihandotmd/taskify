package task

import (
	obModel "github.com/raihandotmd/taskify/internal/adapters/outbound/models/task"
)

type (
	EditTaskRequest struct {
		ID          string `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Status      string `json:"status"`
		Deadline    string `json:"deadline"`
	}

	EditTaskResponse struct {
		ID          string `json:"id"`
		ProjectID   string `json:"project_id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Status      string `json:"status"`
		Deadline    string `json:"deadline"`
	}
)

func (req EditTaskRequest) ToObModel() obModel.Task {
	return obModel.Task{
		ID:          req.ID,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		Deadline:    req.Deadline,
	}
}
