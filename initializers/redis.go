package initializers

import (
	"github.com/go-redis/redis/v8"
)

func RedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "tasikmalaya123..",
		DB:       0,
	})
}
