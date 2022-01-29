package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisConfig struct {
	Password string
}

func NewRedisConfig(key string) *RedisConfig {
	return &RedisConfig{
		Password: key,
	}
}

type RedisClientManager interface {
	Set(key string, value interface{}, exp time.Duration) error
	Get(key string) (string, error)
	HGetAll(key string, dst interface{}) error
}

type RedisClient struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisClient(password string) *RedisClient {
	c := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     ":6379",
		Password: "yourpassword",
		DB:       0,
	})

	return &RedisClient{
		client: rdb,
		ctx:    c,
	}
}

func (r *RedisClient) Set(key string, value interface{}, exp time.Duration) error {
	return r.client.Set(r.ctx, key, value, exp).Err()
}

func (r *RedisClient) Get(key string) (string, error) {
	return r.client.Get(r.ctx, key).Result()
}

func (r *RedisClient) HGetAll(key string, dst interface{}) error {
	if err := r.client.HGetAll(r.ctx, "key").Scan(&dst); err != nil {
		return err
	}
	return nil
}


