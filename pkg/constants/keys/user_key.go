package keys

import "time"

const (
	Users = "users"
	User  = "user"
)

// User module cache key constants (backward compatible).
// Other modules can use keys.AllKey("product"), keys.ItemKeyFormat("product"), etc.
var (
	UserAllKey              = AllKey(Users) // scan pattern for invalidation
	UserKey                 = ItemKeyFormat(User)
	UserPaginationKey       = PaginationKeyFormat(Users)
	UserPaginationSearchKey = PaginationSearchKeyFormat(Users)
	UserListKey             = ListKey(Users)
)

// ExpireTime is the default expiration for user caches (1 day).
const ExpireTime = DefaultExpiration

// UserInvalidationKeys returns Keys for invalidating all user caches.
func UserInvalidationKeys() Keys {
	return InvalidationKeys(Users)
}

// UserInvalidationKeysWithExpiration returns Keys with custom expiration.
func UserInvalidationKeysWithExpiration(expiration time.Duration) Keys {
	return InvalidationKeysWithExpiration(Users, expiration)
}
