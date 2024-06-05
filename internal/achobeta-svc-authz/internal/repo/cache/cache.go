package cache

import (
	"achobeta-svc/internal/achobeta-svc-website/config"
	"context"
	"fmt"

	redis "github.com/redis/go-redis/v9"
)

// Cache is an interface
type Cache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string) error
}

type impl struct {
	cli *redis.Client
}

func New() Cache {
	// 应该由 go lib 提供统一的 New 方法，用于初始化 Redis
	return &impl{
		cli: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", config.Get().Database.Redis.Host, config.Get().Database.Redis.Port),
			Password: config.Get().Database.Redis.Password,
		}),
	}
}

func (i *impl) Get(ctx context.Context, key string) (string, error) {
	return i.cli.Get(ctx, key).Result()
}

func (i *impl) Set(ctx context.Context, key string, value string) error {
	return i.cli.Set(ctx, key, value, 0).Err()
}
