package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var defaultCache Cache

type Cache interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
	Close() error
}

type redisCache struct {
	client *redis.Client
}

func Init(redisURL string) error {
	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		return fmt.Errorf("parsing redis url: %w", err)
	}

	client := redis.NewClient(opts)
	
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("pinging redis: %w", err)
	}

	defaultCache = &redisCache{client: client}
	return nil
}

func SetDefault(c Cache) {
	defaultCache = c
}

func GetDefault() Cache {
	return defaultCache
}

func (r *redisCache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}

func (r *redisCache) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *redisCache) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

func (r *redisCache) Close() error {
	return r.client.Close()
}

// Package-level helpers for default cache
func Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return defaultCache.Set(ctx, key, value, expiration)
}

func Get(ctx context.Context, key string) (string, error) {
	return defaultCache.Get(ctx, key)
}

func Delete(ctx context.Context, key string) error {
	return defaultCache.Delete(ctx, key)
}
