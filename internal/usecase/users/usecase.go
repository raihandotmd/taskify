package users

import (
	obCacheUser "github.com/raihandotmd/taskify/internal/adapters/outbound/repositories/cache/user"
	repositoryUser "github.com/raihandotmd/taskify/internal/adapters/outbound/repositories/sql/user"
)

type (
	UseCase interface {
		IRegister
		ILogin
		ILogout
	}

	ucUser struct {
		register IRegister
		login    ILogin
		logout   ILogout
	}
)

func New() UseCase {
	return &ucUser{}
}

var (
	repoUser  = repositoryUser.New()
	cacheUser = obCacheUser.New()
)
