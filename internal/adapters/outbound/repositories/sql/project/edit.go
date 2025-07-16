package project

import (
	"context"

	obModel "github.com/raihandotmd/taskify/internal/adapters/outbound/models/project"
	"gorm.io/gorm"
)

func (r *repository) EditProject(ctx context.Context, orm *gorm.DB, req obModel.Project) (obModel.Project, error) {

	// update project details in the database
	if err := orm.WithContext(ctx).Model(&obModel.Project{}).Where("id = ? AND created_by = ?", req.ID, req.UserID).Updates(req).Error; err != nil {
		return obModel.Project{}, err
	}

	// retrieve the updated project to return
	var updatedProject obModel.Project
	if err := orm.WithContext(ctx).Where("id = ?", req.ID).First(&updatedProject).Error; err != nil {
		return obModel.Project{}, err
	}
	return updatedProject, nil
}

type IEdit interface {
	EditProject(ctx context.Context, orm *gorm.DB, project obModel.Project) (obModel.Project, error)
}
