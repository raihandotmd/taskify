package project

type (
	Repository interface {
		ICreate
		IGetAllProjectByUserId
		IEdit
		IDelete
	}

	repository struct {
		create                ICreate
		getAllProjectByUserId IGetAllProjectByUserId
		edit                  IEdit
		delete                IDelete
	}
)

func New() Repository {
	return &repository{}
}
