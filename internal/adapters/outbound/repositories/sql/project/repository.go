package project

type (
	Repository interface {
		ICreate
		IGetAllProjectByUserId
	}

	repository struct {
		create                ICreate
		getAllProjectByUserId IGetAllProjectByUserId
	}
)

func New() Repository {
	return &repository{}
}
