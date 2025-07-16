package project

import (
	"context"

	obModel "github.com/raihandotmd/taskify/internal/adapters/outbound/models/project"
	"gorm.io/gorm"
)

func (r *repository) DeleteProject(ctx context.Context, orm *gorm.DB, projectId string, userId string) error {
	if err := orm.WithContext(ctx).Where("id = ? AND created_by = ?", projectId, userId).Delete(&obModel.Project{}).Error; err != nil {
		return err
	}

	return nil
}

type IDelete interface {
	DeleteProject(ctx context.Context, orm *gorm.DB, projectId string, userId string) error
}
