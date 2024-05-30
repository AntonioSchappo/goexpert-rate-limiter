package repository

import (
	"time"

	"github.com/go-redis/redis"
)

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository(addr, password string, db int) *RedisRepository {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return &RedisRepository{client: client}
}

func (r *RedisRepository) Set(key string, value string, timeout time.Duration) error {
	return r.client.Set(key, value, timeout).Err()
}

func (r *RedisRepository) Get(key string) (string, error) {
	return r.client.Get(key).Result()
}

func (r *RedisRepository) Incr(key string) error {
	return r.client.Incr(key).Err()
}
