package redis

import (
	"github.com/go-redis/redis/v8"
	"github.com/patrickmn/go-cache"
)

type SnsInfoRedis struct {
	rds *redis.Client
	c   *cache.Cache
}
