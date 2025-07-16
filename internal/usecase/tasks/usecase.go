package tasks

import (
	repositoryTask "github.com/raihandotmd/taskify/internal/adapters/outbound/repositories/sql/task"
)

var repoTask = repositoryTask.New()

type (
	UseCase interface {
		INewTask
		IGetAllTasksByProjectId
		IEditTask
		IDeleteTask
	}

	ucTask struct {
		newTask                INewTask
		getAllTasksByProjectId IGetAllTasksByProjectId
		editTask               IEditTask
		deleteTask             IDeleteTask
	}
)

func New() UseCase {
	return &ucTask{}
}
