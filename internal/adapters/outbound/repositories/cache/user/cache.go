package user

const (
	serviceKey       = "user-service"
	revokeOneRepoKey = serviceKey + ":revokeOne:%s"
)

type (
	Cache interface {
		IRevokeUserToken
		IExistToken
	}

	cache struct {
		revokeUserToken IRevokeUserToken
		existToken      IExistToken
	}
)

func New() Cache {
	return &cache{}
}
