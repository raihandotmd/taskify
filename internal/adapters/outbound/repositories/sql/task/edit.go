package task

import (
	"context"

	obModel "github.com/raihandotmd/taskify/internal/adapters/outbound/models/task"
	"gorm.io/gorm"
)

func (r *repository) EditTask(ctx context.Context, orm *gorm.DB, task obModel.Task) (obModel.Task, error) {
	var existingTask obModel.Task

	if err := orm.WithContext(ctx).Where("id = ?", task.ID).First(&existingTask).Error; err != nil {
		return obModel.Task{}, err
	}

	if err := orm.WithContext(ctx).Model(&existingTask).Updates(task).Error; err != nil {
		return obModel.Task{}, err
	}

	return existingTask, nil
}

type IEdit interface {
	EditTask(ctx context.Context, orm *gorm.DB, task obModel.Task) (obModel.Task, error)
}
