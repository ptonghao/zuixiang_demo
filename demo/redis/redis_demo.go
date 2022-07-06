/*
 * @Author: Jimpu
 * @Description: redis string
 */

package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

// RedisData .
type RedisData struct {
	// key
	Key string
	// value
	Value interface{}
	// 过期时间, 小时: time.Hour、分钟: time.Minute、秒: time.Second
	DTTL time.Duration
}

// RedisCache redis instance
var RedisInstace = &redis.Client{}

func init() {
	RedisInstace = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "123456",
		DB:       0,
	})

	//ping
	pong, err := RedisInstace.Ping().Result()
	if err != nil {
		fmt.Println("ping error", err.Error())
		return
	}
	fmt.Println("ping result:", pong)
}

// Get 获取缓存数据
func (r *RedisData) Get() (string, error) {
	result, err := RedisInstace.Get(r.Key).Result()
	return result, err
}

// Set 设置数据
func (r *RedisData) Set() error {
	err := RedisInstace.Set(r.Key, r.Value, r.DTTL).Err()
	return err
}
