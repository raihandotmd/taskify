package tasks

import (
	"context"

	ucModel "github.com/raihandotmd/taskify/internal/usecase/models/task"
	taskifyGorm "github.com/raihandotmd/taskify/pkg/gorm"
)

func (uc *ucTask) DeleteTask(ctx context.Context, taskId string) (ucModel.DeleteTaskResponse, error) {
	var orm = taskifyGorm.DB

	// delete task
	if err := repoTask.DeleteTask(ctx, orm, taskId); err != nil {
		return ucModel.DeleteTaskResponse{}, err
	}

	return ucModel.DeleteTaskResponse{
		Message: "Task deleted successfully",
	}, nil
}

type IDeleteTask interface {
	DeleteTask(ctx context.Context, taskId string) (ucModel.DeleteTaskResponse, error)
}
