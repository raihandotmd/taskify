package user

import (
	"context"
	"fmt"

	obModel "github.com/raihandotmd/taskify/internal/adapters/outbound/models/user"
	taskifyRedis "github.com/raihandotmd/taskify/pkg/redis"
)

func (c *cache) ExistToken(ctx context.Context, req obModel.RevokeUserToken) (bool, error) {
	var redisCilent = taskifyRedis.RDB
	redisKey := fmt.Sprintf(revokeOneRepoKey, req.UserID)

	// check if the user token exists in the cache
	exists, err := redisCilent.Exists(ctx, redisKey).Result()
	if err != nil {
		return false, fmt.Errorf("failed to check if user token exists in cache: %v", err)
	}

	return exists != 0, nil
}

type IExistToken interface {
	ExistToken(ctx context.Context, req obModel.RevokeUserToken) (bool, error)
}
