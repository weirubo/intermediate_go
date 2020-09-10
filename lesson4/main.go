package main

// go-redis 包
import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var redisClient *redis.Client
var ctx = context.Background()

func InitRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})
	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("redis 错误信息：%v\n", err)
		return
	}
	fmt.Println(pong)
}

func InitRedis2() {
	options, err := redis.ParseURL("redis://127.0.0.1:6379/0")
	if err != nil {
		fmt.Printf("redis 错误信息:%v\n", err)
		return
	}
	redisClient = redis.NewClient(options)
	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("redis 错误信息：%v\n", err)
		return
	}
	fmt.Println(pong)
}
func main() {
	InitRedis()
	// InitRedis2()
	/*val, err := Set("price", 99, time.Second*300)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)*/

	val, err := Get("price")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
}

// Command
// set
func Set(key string, value interface{}, expiration time.Duration) (val string, err error) {
	val, err = redisClient.Set(ctx, key, value, expiration).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

// get
func Get(key string) (string, error) {
	val, err := redisClient.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
