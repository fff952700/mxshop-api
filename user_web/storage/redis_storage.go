package storage

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	Client     *redis.Client
	Expiration time.Duration
}

// NewRedisStore creates a new RedisStore instance
func NewRedisStore(client *redis.Client, expiration time.Duration) *RedisStore {
	return &RedisStore{
		Client:     client,
		Expiration: expiration,
	}
}

// Set stores the captcha value in Redis
func (r *RedisStore) Set(id string, value string) error {
	return r.Client.Set(context.Background(), id, value, r.Expiration).Err()
}

// Get retrieves the captcha value from Redis
func (r *RedisStore) Get(id string, clear bool) string {
	value, err := r.Client.Get(context.Background(), id).Result()
	if err != nil {
		return ""
	}
	if clear {
		r.Client.Del(context.Background(), id)
	}
	return value
}

// Verify checks if the provided answer matches the stored captcha value
func (r *RedisStore) Verify(id string, answer string, clear bool) bool {
	value := r.Get(id, clear)
	return value == answer
}
