package user

type (
	Repository interface {
		ICreate
		IGetUserByEmail
	}

	repository struct {
		create         ICreate
		getUserByEmail IGetUserByEmail
	}
)

func New() Repository {
	return &repository{}
}
