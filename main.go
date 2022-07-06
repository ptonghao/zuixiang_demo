/*
 * @Author: Jimpu
 * @Description: main
 */

package main

import (
	"Demo1368/demo/josephus"
	"Demo1368/demo/quick_sort"
	"Demo1368/demo/redis"
	"Demo1368/demo/shuffle_cards"
	"fmt"
	"log"
	"time"
)

func main() {
	// 1、实现一个函数，输入为任意长度的int64数组，返回元素最大差值，例如输入arr=[5,8,10,1,3]，返回9。
	data := []int{5, 8, 10, 1, 3}
	quick_sort.QuickSort(data, 0, len(data))
	log.Println(fmt.Sprintf("求slice元素最大差值: src=%v, result=%v", data, data[len(data)-1]-data[0]))

	// 2、实现一个函数，对输入的扑克牌执行洗牌，保证其是均匀分布的，也就是说列表中的每一张扑克牌出现在列表的每一个位置上的概率必须相同。
	srcData := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "小王", "大王"}
	log.Println(fmt.Sprintf("洗牌: result=%s", shuffle_cards.ShuffleCards(srcData, len(srcData))))

	// 3. 设计一个带失效时间的缓存数据结构，key和value都是string，并实现增删改查接口。
	key := "TestKey:"
	cache := redis.NewCache(key, 10*time.Second)
	cache.Set("hello cache!")
	log.Println(fmt.Sprintf("cache: key: %v, result: %s", key, cache.Get()))

	// 4. 约瑟夫环问题
	n, m := 100, 3
	log.Println(fmt.Sprintf("约瑟夫环问题: n=%v, m=%v, result=%v", n, m, josephus.Josephus(n, m)))
}
