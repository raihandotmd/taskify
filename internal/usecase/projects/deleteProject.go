package projects

import (
	"context"

	ucModel "github.com/raihandotmd/taskify/internal/usecase/models/project"
	taskifyGorm "github.com/raihandotmd/taskify/pkg/gorm"
)

func (uc *ucProject) DeleteProject(ctx context.Context, projectId string, userId string) (ucModel.DeleteProjectResponse, error) {
	var orm = taskifyGorm.DB

	// delete project
	if err := repoProject.DeleteProject(ctx, orm, projectId, userId); err != nil {
		return ucModel.DeleteProjectResponse{}, err
	}

	return ucModel.DeleteProjectResponse{
		Message: "Project deleted successfully",
	}, nil
}

type IDeleteProject interface {
	DeleteProject(ctx context.Context, projectId string, userId string) (ucModel.DeleteProjectResponse, error)
}
