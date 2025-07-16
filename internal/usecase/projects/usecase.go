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
		IDeleteProject
	}

	ucProject struct {
		newProject            INewProject
		getAllProjectByUserId IGetAllProjectByUserId
		editProject           IEditProject
		deleteProject         IDeleteProject
	}
)

func New() UseCase {
	return &ucProject{}
}
