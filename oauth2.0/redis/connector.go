package redis

import (
	"github.com/go-redis/redis/v8"
	"time"
)

const SnsDB = 1

func Connect(host string, db int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:         host,
		DB:           db,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		MinIdleConns: 2,
		IdleTimeout:  5 * time.Second,
	})
}
