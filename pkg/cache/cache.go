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
