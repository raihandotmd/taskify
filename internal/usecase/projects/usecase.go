package projects

import (
	repositoryProject "github.com/raihandotmd/taskify/internal/adapters/outbound/repositories/sql/project"
)

var repoProject = repositoryProject.New()

type (
	UseCase interface {
		INewProject
		IGetAllProjectByUserId
		IEditProject
	}

	ucProject struct {
		newProject            INewProject
		getAllProjectByUserId IGetAllProjectByUserId
		editProject           IEditProject
	}
)

func New() UseCase {
	return &ucProject{}
}
