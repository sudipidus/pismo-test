package services

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type Lock interface {
	Lock(id string) error
	Unlock(id string) error
}

type RedisLock struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisLock(client *redis.Client) *RedisLock {
	return &RedisLock{
		client: client,
		ctx:    context.Background(),
	}
}

func (r *RedisLock) Lock(id string) error {

	_, err := r.client.SetNX(r.ctx, id, "locked", 0).Result()
	if err != nil {
		return err
	}
	return nil

}

func (r *RedisLock) Unlock(id string) error {
	_, err := r.client.Del(r.ctx, id).Result()
	return err
}
