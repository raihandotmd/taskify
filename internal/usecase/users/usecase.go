package users

type (
	UseCase interface {
		IRegister
	}

	ucUser struct {
		register IRegister
	}
)

func New() UseCase {
	return &ucUser{}
}
