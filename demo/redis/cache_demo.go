/*
 * @Author: Jimpu
 * @Description: 设计一个带失效时间的缓存数据结构，key和value都是string，并实现增删改查接口
 */

package redis

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

// CacheData .
type CacheData struct {
	Cache *cache.Cache
	// key
	Key string
	// value
	Value interface{}
	// 过期时间, 小时: time.Hour、分钟: time.Minute、秒: time.Second
	DTTL time.Duration
}

func NewCache(key string, ttl time.Duration) *CacheData {
	r := &CacheData{
		DTTL:  ttl,
		Key:   key,
		Cache: cache.New(ttl, 1*time.Minute), // 默认过期时间为1day，每1min清理一次过期缓存
	}
	if r.DTTL == 0 {
		r.DTTL = time.Hour * 24
	}
	return r
}

// Get 获取缓存数据
func (r *CacheData) Get() string {
	result, _ := r.Cache.Get(r.Key)
	return fmt.Sprintf("%v", result)
}

// Set 设置数据
func (r *CacheData) Set(value string) {
	if value != "" {
		r.Cache.Set(r.Key, value, r.DTTL)
		return
	}
	r.Cache.Set(r.Key, r.Value, r.DTTL)
}
