package cache

import "errors"

var (
	errCacheMiss = errors.New("cache miss")
)

// IsCacheMiss returns true if the given error is the result of a cache miss.
func IsCacheMiss(err error) bool {
	return errors.Is(err, errCacheMiss)
}
