package database

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var Cache *redis.Client
var CacheChannel chan string

func SetupRedis() {
	Cache = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB: 0,
	})
}

func ClearCache(keys ...string) {
	for _, key := range keys {
	CacheChannel <- key
	}
}

func SetCache(bytes []byte) {
	var ctx = context.Background()
	Cache.Set(ctx, "products_backend", bytes, 60 * time.Minute)
}