package user

type (
	Repository interface {
		ICreate
	}

	repository struct {
		create ICreate
	}
)

func New() Repository {
	return &repository{}
}
