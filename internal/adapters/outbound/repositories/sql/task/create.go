package task

import (
	"context"

	obModel "github.com/raihandotmd/taskify/internal/adapters/outbound/models/task"
	"gorm.io/gorm"
)

func (r *repository) CreateTask(ctx context.Context, orm *gorm.DB, task obModel.Task) (obModel.Task, error) {
	if err := orm.WithContext(ctx).Create(&task).Error; err != nil {
		return obModel.Task{}, err
	}

	return task, nil
}

type ICreate interface {
	CreateTask(ctx context.Context, orm *gorm.DB, task obModel.Task) (obModel.Task, error)
}
