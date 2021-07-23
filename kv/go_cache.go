package kv

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var (
	c            *cache.Cache
	cacheFile    string
	NoExpiration = cache.NoExpiration
)

func init() {
	c = cache.New(5*time.Minute, 10*time.Minute)
}

// Set 设值
func Set(k string, x interface{}) {
	c.Set(k, x, NoExpiration)
}

// Get 取值
func Get(k string) (value interface{}, found bool) {
	value, found = c.Get(k)
	return
}
