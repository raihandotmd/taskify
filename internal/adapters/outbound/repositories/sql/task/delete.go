package task

import (
	"context"

	obModel "github.com/raihandotmd/taskify/internal/adapters/outbound/models/task"
	"gorm.io/gorm"
)

func (r *repository) DeleteTask(ctx context.Context, orm *gorm.DB, taskId string) error {
	if err := orm.WithContext(ctx).Where("id = ?", taskId).Delete(&obModel.Task{}).Error; err != nil {
		return err
	}

	return nil
}

type IDelete interface {
	DeleteTask(ctx context.Context, orm *gorm.DB, taskId string) error
}
