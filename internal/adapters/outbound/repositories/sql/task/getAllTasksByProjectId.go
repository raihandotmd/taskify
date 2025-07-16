package task

import (
	"context"

	obModel "github.com/raihandotmd/taskify/internal/adapters/outbound/models/task"
	"gorm.io/gorm"
)

func (r *repository) GetAllTasksByProjectId(ctx context.Context, orm *gorm.DB, projectId string) ([]obModel.Task, error) {
	var tasks []obModel.Task

	if err := orm.WithContext(ctx).Where("project_id = ?", projectId).Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

type IGetAllTasksByProjectId interface {
	GetAllTasksByProjectId(ctx context.Context, orm *gorm.DB, projectId string) ([]obModel.Task, error)
}
