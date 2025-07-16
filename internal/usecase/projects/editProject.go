package projects

import (
	"context"

	obModel "github.com/raihandotmd/taskify/internal/adapters/outbound/models/project"
	ucModel "github.com/raihandotmd/taskify/internal/usecase/models/project"
	taskifyGorm "github.com/raihandotmd/taskify/pkg/gorm"
)

func (uc *ucProject) EditProject(ctx context.Context, req ucModel.EditProjectRequest, userId string) (ucModel.EditProjectResponse, error) {
	var orm = taskifyGorm.DB

	var (
		project obModel.Project
		err     error
	)

	// update project
	if project, err = repoProject.EditProject(ctx, orm, req.ToObModel(userId)); err != nil {
		return ucModel.EditProjectResponse{}, err
	}

	return ucModel.EditProjectResponse{
		ID:          project.ID,
		Name:        project.Name,
		Description: project.Description,
		UserID:      userId,
	}, nil
}

type IEditProject interface {
	EditProject(ctx context.Context, req ucModel.EditProjectRequest, userId string) (ucModel.EditProjectResponse, error)
}
