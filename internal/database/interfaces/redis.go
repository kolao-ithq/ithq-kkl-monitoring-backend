package interfaces

import (
	"codelab/go-gin-template/pkg/constants/keys"
	"context"
	"encoding/json"
	"time"
)

type RedisRepository interface {
	Get(ctx context.Context, key, path string) (json.RawMessage, error)
	Set(ctx context.Context, key string, path string, value interface{}, expiration time.Duration) error
	Delete(ctx context.Context, key string) error
	Exists(ctx context.Context, key string) (bool, error)
	Increment(ctx context.Context, key string) (int64, error)
	Expire(ctx context.Context, key string, expiration time.Duration) error
	DeleteByPrefix(ctx context.Context, prefixes []string) error
	DeleteAll(ctx context.Context) error
	InvalidateEntityCaches(ctx context.Context, keys keys.Keys)
	Close() error
}
