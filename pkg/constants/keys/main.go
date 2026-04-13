package keys

import (
	"fmt"
	"time"
)

const (
	MyApp   = "my_app:" // app prefix to avoid key collisions
)

// DefaultExpiration is the default cache TTL for entity keys.
const DefaultExpiration = 1 * 24 * time.Hour // 1 day

// Keys is used by cache adapters (e.g. Redis) for invalidation and options.
type Keys struct {
	Key             string
	ItemKeyTemplate string
	EntityID        string
	EntityType      string
	Prefix          []string
	Attempts         int
	RetryDelay      time.Duration
	Expiration      time.Duration
}

// AllKey returns the scan/invalidate pattern for a module (e.g. "my_app:users*").
func AllKey(module string) string {
	return MyApp + module + "*"
}

// ItemKeyFormat returns the format string for a single item key (e.g. "my_app:user:%d").
// Use with fmt.Sprintf(keys.ItemKeyFormat(module), id).
func ItemKeyFormat(module string) string {
	return MyApp + module + ":%d"
}

// ItemKey returns the cache key for one item by id.
func ItemKey(module string, id uint) string {
	return fmt.Sprintf(ItemKeyFormat(module), id)
}

// PaginationKeyFormat returns the format string for pagination (e.g. "my_app:users:pagination:%d").
func PaginationKeyFormat(module string) string {
	return MyApp + module + ":pagination:%d"
}

// PaginationSearchKeyFormat returns the format string for pagination with search.
// Use with fmt.Sprintf(keys.PaginationSearchKeyFormat(module), page, search).
func PaginationSearchKeyFormat(module string) string {
	return MyApp + module + ":pagination:%d:search:%s"
}

// ListKey returns the key for the full list of a module (e.g. "my_app:users:list").
func ListKey(module string) string {
	return MyApp + module + ":list"
}

// InvalidationKeys returns a Keys struct with Key set for InvalidateEntityCaches.
func InvalidationKeys(module string) Keys {
	return Keys{Key: AllKey(module)}
}

// InvalidationKeysWithExpiration returns Keys with Key and Expiration set.
func InvalidationKeysWithExpiration(module string, expiration time.Duration) Keys {
	return Keys{Key: AllKey(module), Expiration: expiration}
}
