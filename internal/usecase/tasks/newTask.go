package tasks

import (
	"context"

	obModel "github.com/raihandotmd/taskify/internal/adapters/outbound/models/task"
	ucModel "github.com/raihandotmd/taskify/internal/usecase/models/task"
	taskifyGorm "github.com/raihandotmd/taskify/pkg/gorm"
)

func (uc *ucTask) NewTask(ctx context.Context, req ucModel.NewTaskRequest) (ucModel.NewTaskResponse, error) {
	var orm = taskifyGorm.DB

	var (
		task obModel.Task
		err  error
	)

	// create task
	if task, err = repoTask.CreateTask(ctx, orm, req.ToObModel()); err != nil {
		return ucModel.NewTaskResponse{}, err
	}

	return ucModel.NewTaskResponse{
		ID:          task.ID,
		ProjectID:   task.ProjectID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		Deadline:    task.Deadline,
	}, nil
}

type INewTask interface {
	NewTask(ctx context.Context, req ucModel.NewTaskRequest) (ucModel.NewTaskResponse, error)
}
