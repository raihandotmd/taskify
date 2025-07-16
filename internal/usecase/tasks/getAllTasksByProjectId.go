package tasks

import (
	"context"

	ucModel "github.com/raihandotmd/taskify/internal/usecase/models/task"
	taskifyGorm "github.com/raihandotmd/taskify/pkg/gorm"
)

func (uc *ucTask) GetAllTasksByProjectId(ctx context.Context, projectId string) (ucModel.GetAllTasksByProjectIdResponse, error) {
	var orm = taskifyGorm.DB

	tasks, err := repoTask.GetAllTasksByProjectId(ctx, orm, projectId)
	if err != nil {
		return ucModel.GetAllTasksByProjectIdResponse{}, err
	}

	return ucModel.GetAllTasksByProjectIdResponse{
		Tasks: tasks,
	}, nil
}

type IGetAllTasksByProjectId interface {
	GetAllTasksByProjectId(ctx context.Context, projectId string) (ucModel.GetAllTasksByProjectIdResponse, error)
}
