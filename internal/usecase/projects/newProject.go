package projects

import (
	"context"

	obModel "github.com/raihandotmd/taskify/internal/adapters/outbound/models/project"
	ucModel "github.com/raihandotmd/taskify/internal/usecase/models/project"
	taskifyGorm "github.com/raihandotmd/taskify/pkg/gorm"
)

func (uc *ucProject) NewProject(ctx context.Context, req ucModel.NewProjectRequest, userId string) (ucModel.NewProjectResponse, error) {
	var orm = taskifyGorm.DB

	var (
		project obModel.Project
		err     error
	)

	// create project
	if project, err = repoProject.CreateProject(ctx, orm, req.ToObModel(userId)); err != nil {
		return ucModel.NewProjectResponse{}, err
	}

	return ucModel.NewProjectResponse{
		ID:          project.ID,
		Name:        project.Name,
		Description: project.Description,
		UserID:      userId,
	}, nil
}

type INewProject interface {
	NewProject(ctx context.Context, req ucModel.NewProjectRequest, userId string) (ucModel.NewProjectResponse, error)
}
