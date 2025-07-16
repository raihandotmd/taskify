package projects

import (
	"context"

	ucModel "github.com/raihandotmd/taskify/internal/usecase/models/project"
	taskifyGorm "github.com/raihandotmd/taskify/pkg/gorm"
)

func (uc *ucProject) GetAllProjectByUserId(ctx context.Context, userId string) (ucModel.GetAllProjectByUserIdResponse, error) {
	var orm = taskifyGorm.DB

	projects, err := repoProject.GetAllProjectByUserId(ctx, orm, userId)
	if err != nil {
		return ucModel.GetAllProjectByUserIdResponse{}, err
	}

	return ucModel.GetAllProjectByUserIdResponse{
		Projects: projects,
	}, nil
}

type IGetAllProjectByUserId interface {
	GetAllProjectByUserId(ctx context.Context, userId string) (ucModel.GetAllProjectByUserIdResponse, error)
}
