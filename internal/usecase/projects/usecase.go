package projects

import (
	repositoryProject "github.com/raihandotmd/taskify/internal/adapters/outbound/repositories/sql/project"
)

var repoProject = repositoryProject.New()

type (
	UseCase interface {
		INewProject
	}

	ucProject struct {
		newProject INewProject
	}
)

func New() UseCase {
	return &ucProject{}
}
