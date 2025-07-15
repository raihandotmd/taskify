package usecase

import (
	"github.com/raihandotmd/taskify/internal/usecase/token"
	"github.com/raihandotmd/taskify/internal/usecase/users"
)

type (
	UseCase interface {
		token.UseCase
		users.UseCase
	}

	ucImpl struct {
		token token.UseCase
		users users.UseCase
	}
)
