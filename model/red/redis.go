package red

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

//red.SetKey("zz","ssf",1*time.Minute)
//val,err := red.GetKey("ji")
//if err {
//	fmt.Print(val)
//}else {
//	fmt.Print("df")
//}
//time.Sleep(time.Second * 1) // 暂停一秒等待 subscribe 完成
////emqx.Publish(mqt,0)
//fmt.Println(utils.GetUUID())

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
		panic(any(err))
	}
}

func GetKey(key string) (string, bool) {
	ctx := context.Background()
	val, err := LinkRedis().Get(ctx, key).Result()
	if err == redis.Nil {
		return "", false
	} else if err != nil {
		panic(any(err))
	} else {
		return val, true
	}
}
