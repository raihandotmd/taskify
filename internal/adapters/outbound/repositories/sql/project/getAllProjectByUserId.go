package project

import (
	"context"

	obModel "github.com/raihandotmd/taskify/internal/adapters/outbound/models/project"
	"gorm.io/gorm"
)

func (r *repository) GetAllProjectByUserId(ctx context.Context, orm *gorm.DB, userId string) ([]obModel.Project, error) {
	var projects []obModel.Project
	if err := orm.WithContext(ctx).Where("user_id = ?", userId).Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

type IGetAllProjectByUserId interface {
	GetAllProjectByUserId(ctx context.Context, orm *gorm.DB, userId string) ([]obModel.Project, error)
}
