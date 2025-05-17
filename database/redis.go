package database

import (
	"context"
	"golang-rest-api-template/env"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value any, expiration time.Duration) *redis.StatusCmd
	Keys(context.Context, string) *redis.StringSliceCmd
	Del(context.Context, ...string) *redis.IntCmd
}

func NewRedisClient(env *env.Env) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:            env.RedisAddress,
		Password:        env.RedisPassword,
		DB:              0,
		MaxIdleConns:    env.RedisMaxIdle,
		MaxActiveConns:  env.RedisMaxActive,
		ConnMaxIdleTime: env.RedisIdleTimeout,
	})
}
