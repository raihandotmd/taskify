package project

import (
	"context"

	obModel "github.com/raihandotmd/taskify/internal/adapters/outbound/models/project"
	"gorm.io/gorm"
)

func (r *repository) CreateProject(ctx context.Context, orm *gorm.DB, req obModel.Project) (obModel.Project, error) {
	if err := orm.WithContext(ctx).Create(&req).Error; err != nil {
		return obModel.Project{}, err
	}
	return req, nil
}

type ICreate interface {
	CreateProject(ctx context.Context, orm *gorm.DB, project obModel.Project) (obModel.Project, error)
}
