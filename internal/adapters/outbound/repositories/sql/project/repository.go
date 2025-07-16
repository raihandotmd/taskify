package project

type (
	Repository interface {
		ICreate
		IGetAllProjectByUserId
		IEdit
	}

	repository struct {
		create                ICreate
		getAllProjectByUserId IGetAllProjectByUserId
		edit                  IEdit
	}
)

func New() Repository {
	return &repository{}
}
