package database

import (
	"context"
	"golang-rest-api-template/env"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis interface {
	Get(key string) *redis.StringCmd
	Set(key string, value any, expiration time.Duration) *redis.StatusCmd
	Del(...string) *redis.IntCmd
	XAdd(a *redis.XAddArgs) *redis.StringCmd
	XReadGroup(a *redis.XReadGroupArgs) *redis.XStreamSliceCmd
	XGroupCreateMkStream(stream, group, start string) *redis.StatusCmd
	XAck(stream, group string, ids ...string) *redis.IntCmd
	Close() error
}

type redisClient struct {
	Client *redis.Client
	Ctx    context.Context
}

func NewRedisClient(ctx context.Context, env *env.Env) Redis {
	return &redisClient{
		Client: redis.NewClient(&redis.Options{
			Addr:            env.RedisAddress,
			Password:        env.RedisPassword,
			DB:              0,
			MaxIdleConns:    env.RedisMaxIdle,
			MaxActiveConns:  env.RedisMaxActive,
			ConnMaxIdleTime: env.RedisIdleTimeout,
		}),
		Ctx: ctx,
	}
}

func (c *redisClient) Get(key string) *redis.StringCmd {
	return c.Client.Get(c.Ctx, key)
}

func (c *redisClient) Set(key string, value any, expiration time.Duration) *redis.StatusCmd {
	return c.Client.Set(c.Ctx, key, value, expiration)
}

func (c *redisClient) Del(key ...string) *redis.IntCmd {
	return c.Client.Del(c.Ctx, key...)
}

func (c *redisClient) XAck(stream, group string, ids ...string) *redis.IntCmd {
	return c.Client.XAck(c.Ctx, stream, group, ids...)
}

func (c *redisClient) XAdd(a *redis.XAddArgs) *redis.StringCmd {
	return c.Client.XAdd(c.Ctx, a)
}

func (c *redisClient) XReadGroup(a *redis.XReadGroupArgs) *redis.XStreamSliceCmd {
	return c.Client.XReadGroup(c.Ctx, a)
}

func (c *redisClient) XGroupCreateMkStream(stream, group, start string) *redis.StatusCmd {
	return c.Client.XGroupCreateMkStream(c.Ctx, stream, group, start)
}

func (c *redisClient) Close() error {
	return c.Client.Close()
}
