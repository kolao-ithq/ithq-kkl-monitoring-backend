package redis

import (
	"codelab/go-gin-template/internal/database/interfaces"
	"codelab/go-gin-template/pkg/constants/keys"
	logs "codelab/go-gin-template/pkg/logger"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// Config holds the configuration for Redis
type Config struct {
	Host         string
	Port         string
	Password     string
	DB           int
	ClearOnStart bool
}

type CacheRepository struct {
	client *redis.Client
	pubsub *redis.PubSub
}

// Get retrieves a value from the cache by key
func (r *CacheRepository) Get(ctx context.Context, key string, path string) (json.RawMessage, error) {
	raw, err := r.client.JSONGet(ctx, key, path).Result()
	if err != nil {
		logs.Error(fmt.Sprintf("failed to get data from cache: %v", err))
		return nil, err
	}
	return json.RawMessage(raw), nil
}

// Set stores a value in the cache with an optional expiration time
func (r *CacheRepository) Set(ctx context.Context, key, path string, value interface{}, expiration time.Duration) error {
	err := r.client.JSONSet(ctx, key, path, value).Err()
	if err != nil {
		logs.Error(fmt.Sprintf("failed to set data in cache: %v", err))
	}

	if expiration > 0 {
		logs.Info(fmt.Sprintf("setting expiration time for key: %s", key))
		return r.client.Expire(ctx, key, expiration).Err()
	}
	return nil
}

// Delete removes a value from the cache by key
func (r *CacheRepository) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

// Exists checks if a key exists in the cache
func (r *CacheRepository) Exists(ctx context.Context, key string) (bool, error) {
	result, err := r.client.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return result > 0, nil
}

// Increment increments a numeric value in the cache
func (r *CacheRepository) Increment(ctx context.Context, key string) (int64, error) {
	return r.client.Incr(ctx, key).Result()
}

// Expire sets an expiration time on a key
func (r *CacheRepository) Expire(ctx context.Context, key string, expiration time.Duration) error {
	return r.client.Expire(ctx, key, expiration).Err()
}

// DeleteAll implements output.CacheRepository.
func (r *CacheRepository) DeleteAll(ctx context.Context) error {
	return r.client.FlushAll(ctx).Err()
}

// DeleteByPrefix implements output.CacheRepository.
func (r *CacheRepository) DeleteByPrefix(ctx context.Context, prefixes []string) error {
	for _, prefix := range prefixes {
		iter := r.client.Scan(ctx, 0, prefix+"*", 0).Iterator()
		for iter.Next(ctx) {
			if err := r.client.Del(ctx, iter.Val()).Err(); err != nil {
				return err
			}
		}
		if err := iter.Err(); err != nil {
			return err
		}
	}
	return nil
}

func ClearCacheOnStart(r *CacheRepository, ctx context.Context, AppKey string) error {
	keys, err := r.client.Keys(ctx, AppKey+"*").Result()
	if err != nil {
		logs.Error(fmt.Sprintf("failed to get keys: %v", err))
		return err
	}
	for _, key := range keys {
		if err := r.Delete(ctx, key); err != nil {
			logs.Error(fmt.Sprintf("failed to delete key %s: %v", key, err))
			return err
		}
	}
	return nil
}

// InvalidateEntityCaches invalidates both list and item caches for any entity
// Uses generics for type safety while supporting different ID types
func (r *CacheRepository) InvalidateEntityCaches(ctx context.Context, key keys.Keys) {
	// Get all keys starting with my_app:users
	keys, err := r.client.Keys(ctx, key.Key).Result()
	if err != nil {
		logs.Error(fmt.Sprintf("failed to get user keys: %v", err))
	}

	if len(keys) == 0 {
		logs.Info("no user keys to delete")
	}
	logs.Info(fmt.Sprintf("keys: %v", keys))
	// Delete all keys found
	for _, k := range keys {
		if err := r.Delete(ctx, k); err != nil {
			logs.Error(fmt.Sprintf("failed to delete key %s: %v", k, err))
		}
	}
}

// Close closes the Redis client
func (r *CacheRepository) Close() error {
	if r.pubsub != nil {
		r.pubsub.Close()
	}

	return r.client.Close()
}

func NewRedisRepository(config Config) (interfaces.RedisRepository, error) {
	// Create Redis client with reconnection and retry options
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Host, config.Port),
		Password: config.Password,
		DB:       config.DB,
	})

	// Test the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := rdb.Ping(ctx).Result()
	if err != nil {
		logs.Error(fmt.Sprintf("can not connect to redis: %v", err))
		return nil, err
	}

	cr := &CacheRepository{client: rdb}

	// Only clear cache if configured to do so
	if config.ClearOnStart {
		clearCtx, clearCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer clearCancel()

		// Clear only your app's cache keys
		prefixes := keys.MyApp
		if err := ClearCacheOnStart(cr, clearCtx, prefixes); err != nil {
			logs.Error(fmt.Sprintf("failed to clear cache prefix %s: %v", prefixes, err))
		} else {
			logs.Info(fmt.Sprintf("cleared cache with prefix: %s", prefixes))
		}
	}

	logs.Info(fmt.Sprintf("redis connection success: %s", result))
	return cr, nil
}

// var _ output.RedisRepository = (*CacheRepository)(nil)
