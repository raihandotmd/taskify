package task

type (
	Repository interface {
		ICreate
		IGetAllTasksByProjectId
		IEdit
		IDelete
	}

	repository struct {
		create                 ICreate
		getAllTasksByProjectId IGetAllTasksByProjectId
		edit                   IEdit
		delete                 IDelete
	}
)

func New() Repository {
	return &repository{}
}
