package cache

import (
	"blackhole-blog/models/dto"
	"github.com/karlseguin/ccache/v3"
	"time"
)

var (
	User    = ccache.New(ccache.Configure[dto.UserDto]())
	Article = ccache.New(ccache.Configure[dto.ArticleDto]())
)

// DeferredSet is a helper function to set cache after function return
func DeferredSet[T any](cache *ccache.Cache[T], key string, item *T, err *error) func() {
	return func() {
		if *err == nil {
			cache.Set(key, *item, 30*time.Minute)
		}
	}
}

// DeferredSetWithRevocer is a helper function to set cache after function return.
// use with panic recover
func DeferredSetWithRevocer[T any](cache *ccache.Cache[T], key string, item *T) func() {
	return func() {
		if r := recover(); r != nil {
			panic(r)
		}
		cache.Set(key, *item, 30*time.Minute)
	}
}

// DeferredDelete is a helper function to delete cache after function return
func DeferredDelete[T any](cache *ccache.Cache[T], key string, err *error) func() {
	return func() {
		if *err == nil {
			cache.Delete(key)
		}
	}
}

// DeferredDeleteWithRevocer is a helper function to delete cache after function return.
// use with panic recover
func DeferredDeleteWithRevocer[T any](cache *ccache.Cache[T], key string) func() {
	return func() {
		if r := recover(); r != nil {
			panic(r)
		}
		cache.Delete(key)
	}
}
