package tasks

import (
	"context"

	obModel "github.com/raihandotmd/taskify/internal/adapters/outbound/models/task"
	ucModel "github.com/raihandotmd/taskify/internal/usecase/models/task"
	taskifyGorm "github.com/raihandotmd/taskify/pkg/gorm"
)

func (uc *ucTask) EditTask(ctx context.Context, req ucModel.EditTaskRequest) (ucModel.EditTaskResponse, error) {
	var orm = taskifyGorm.DB

	var (
		task obModel.Task
		err  error
	)

	// update task
	if task, err = repoTask.EditTask(ctx, orm, req.ToObModel()); err != nil {
		return ucModel.EditTaskResponse{}, err
	}

	return ucModel.EditTaskResponse{
		ID:          task.ID,
		ProjectID:   task.ProjectID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		Deadline:    task.Deadline,
	}, nil
}

type IEditTask interface {
	EditTask(ctx context.Context, req ucModel.EditTaskRequest) (ucModel.EditTaskResponse, error)
}
