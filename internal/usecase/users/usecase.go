package users

import repositoryUser "github.com/raihandotmd/taskify/internal/adapters/outbound/repositories/sql/user"

type (
	UseCase interface {
		IRegister
		ILogin
	}

	ucUser struct {
		register IRegister
		login    ILogin
	}
)

func New() UseCase {
	return &ucUser{}
}

var repoUser = repositoryUser.New()
