package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"server/internal/core/config"
)

func NewRedis(cfg *config.Redis) (*redis.Client, error) {
	if cfg == nil || cfg.Addr == "" {
		fmt.Println("\033[33m[WARN] redis disable\033[0m")
		return nil, nil
	}
	rdb := redis.NewClient(cfg.Options())

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("\033[31m%s\033[0m\n", fmt.Errorf("redis 连接失败: %w", err))
	}

	return rdb, nil
}
