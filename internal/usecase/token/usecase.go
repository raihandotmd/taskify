package token

type (
	UseCase interface {
		IGenerate
	}

	ucAuthToken struct {
		generate IGenerate
	}
)

func New() UseCase {
	return &ucAuthToken{}
}
