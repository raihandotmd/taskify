package taskifyRedis

import (
	"fmt"

	taskifyConfig "github.com/raihandotmd/taskify/internal/config"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
)

var (
	RDB *redis.Client
	ctx = context.Background()
)

func ConnectCache() error {
	redisConfig, err := taskifyConfig.NewRedisConfig()
	if err != nil {
		return fmt.Errorf("failed to load Redis configuration: %v", err)
	}
	RDB = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})
	return nil
}
