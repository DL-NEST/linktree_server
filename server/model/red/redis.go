package red

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

func LinkRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}

func SetKey(key string, value string, time time.Duration) {
	ctx := context.Background()
	err := LinkRedis().Set(ctx, key, value, time).Err()
	if err != nil {
		panic(err)
	}
}

func GetKey(key string) (string, bool) {
	ctx := context.Background()
	val, err := LinkRedis().Get(ctx, key).Result()
	if err == redis.Nil {
		return "", false
	} else if err != nil {
		panic(err)
	} else {
		return val, true
	}
}
