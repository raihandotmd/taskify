package task

import (
	obModel "github.com/raihandotmd/taskify/internal/adapters/outbound/models/task"
)

type (
	NewTaskRequest struct {
		ProjectID   string `json:"project_id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Status      string `json:"status"`
		Deadline    string `json:"deadline"`
	}

	NewTaskResponse struct {
		ID          string `json:"id"`
		ProjectID   string `json:"project_id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Status      string `json:"status"`
		Deadline    string `json:"deadline"`
	}
)

func (req NewTaskRequest) ToObModel() obModel.Task {
	status := req.Status
	if status == "" {
		status = "todo"
	}

	return obModel.Task{
		ProjectID:   req.ProjectID,
		Title:       req.Title,
		Description: req.Description,
		Status:      status,
		Deadline:    req.Deadline,
	}
}
