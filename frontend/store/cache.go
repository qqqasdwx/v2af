package store

import (
	"log"
	"time"

	"horus/common/models"
	"horus/frontend/config"
	"github.com/toolkits/cache"
)

func InitCache() {
	c := config.Config()
	if c.Cache.Provider == "memory" {
		cache.Instance = cache.NewInMemoryCache(c.Cache.Expire)
	} else if c.Cache.Provider == "redis" {
		cache.Instance = cache.NewRedisCache(
			c.Redis.Addr,
			c.Redis.Idle,
			c.Redis.Max,
			time.Duration(c.Redis.Timeout.Conn)*time.Millisecond,
			time.Duration(c.Redis.Timeout.Read)*time.Millisecond,
			time.Duration(c.Redis.Timeout.Write)*time.Millisecond,
			c.Cache.Expire,
		)
	} else {
		log.Fatalln("cache.provider should be memory or redis")
	}

	models.CachePrefix = c.Cache.Prefix
	// 有些不影响水平扩展的东西总是适合缓存在内存里
	models.MemoryCache = cache.NewInMemoryCache(c.Cache.Expire)
}
