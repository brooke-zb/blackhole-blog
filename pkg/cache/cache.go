package cache

import (
	"blackhole-blog/models/dto"
	"github.com/karlseguin/ccache/v3"
	"time"
)

var User = ccache.New(ccache.Configure[dto.UserDto]())

// DeferredSetCache is a helper function to set cache after function return
func DeferredSetCache[T any](cache *ccache.Cache[T], key string, item *T, err *error) func() {
	return func() {
		if *err == nil {
			cache.Set(key, *item, 30*time.Minute)
		}
	}
}

// DeferredSetCacheWithRevocer is a helper function to set cache after function return.
// use with panic recover
func DeferredSetCacheWithRevocer[T any](cache *ccache.Cache[T], key string, item *T) func() {
	return func() {
		if r := recover(); r != nil {
			panic(r)
		}
		cache.Set(key, *item, 30*time.Minute)
	}
}
