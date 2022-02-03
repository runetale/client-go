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
	HSet(key string, value interface{}) error
	Get(key string) (string, error)
	HGet(key string, field string) ([]byte, error)
	HGetAll(key string, dst interface{}) error
}

type RedisClient struct {
	Client *redis.Client
	Ctx    context.Context
}

func NewRedisClient(password string) *RedisClient {
	c := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     ":6379",
		Password: password,
		DB:       0,
	})

	return &RedisClient{
		Client: rdb,
		Ctx:    c,
	}
}

func (r *RedisClient) Tx(commandList func() error) error {
	return r.Client.Watch(r.Ctx, func(tx *redis.Tx) error {
		_, err := tx.TxPipelined(r.Ctx, func(pipe redis.Pipeliner) error {
			return commandList()
		})
		return err
	})
}

func (r *RedisClient) Exists(key string) (int64, error) {
	return r.Client.Exists(r.Ctx, key).Result()
}

func (r *RedisClient) Set(key string, value interface{}, exp time.Duration) error {
	return r.Client.Set(r.Ctx, key, value, exp).Err()
}

func (r *RedisClient) HSet(key string, value interface{}) error {
	return r.Client.HSet(r.Ctx, key, value).Err()
}

func (r *RedisClient) Get(key string) ([]byte, error) {
	return r.Client.Get(r.Ctx, key).Bytes()
}

func (r *RedisClient) HGet(key string, field string) ([]byte, error) {
	return r.Client.HGet(r.Ctx, key, field).Bytes()
}

func (r *RedisClient) HGetAll(key string, dst interface{}) error {
	if err := r.Client.HGetAll(r.Ctx, key).Scan(&dst); err != nil {
		return err
	}
	return nil
}
