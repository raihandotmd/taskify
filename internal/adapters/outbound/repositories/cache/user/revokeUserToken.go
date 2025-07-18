package user

import (
	"context"
	"fmt"

	obModel "github.com/raihandotmd/taskify/internal/adapters/outbound/models/user"
	taskifyRedis "github.com/raihandotmd/taskify/pkg/redis"
)

func (c *cache) RevokeUserToken(ctx context.Context, req obModel.RevokeUserToken) error {
	var redisCilent = taskifyRedis.RDB
	redisKey := fmt.Sprintf(revokeOneRepoKey, req.UserID)

	if err := redisCilent.Set(ctx, redisKey, req.UserID, req.ExpiresAt).Err(); err != nil {
		return fmt.Errorf("failed to revoke user token in cache: %v", err)
	}

	return nil
}

type IRevokeUserToken interface {
	RevokeUserToken(ctx context.Context, req obModel.RevokeUserToken) error
}
