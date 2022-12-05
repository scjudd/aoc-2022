package cache

import "errors"

var (
	errCacheMiss = errors.New("cache miss")
)

func IsCacheMiss(err error) bool {
	return errors.Is(err, errCacheMiss)
}
