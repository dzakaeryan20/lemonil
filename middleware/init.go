package middleware

import (
	cachedal "github.com/dzakaeryan20/lemonilo/dal/cache"
)

type Middleware struct {
	cacheDAL cachedal.CacheDAL
}

func New(cache cachedal.CacheDAL) *Middleware {
	return &Middleware{
		cacheDAL: cache,
	}
}
