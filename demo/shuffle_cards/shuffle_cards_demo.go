/*
 * @Author: Jimpu
 * @Description: 洗牌 demo
 */

package shuffle_cards

import (
	"math"
	"math/rand"
	"time"
)

// getRand 生成随机数
func getRand() int {
	min, max := 1, math.MaxInt
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

// ShuffleCards 洗牌
func ShuffleCards(data []string, l int) []string {
	tmpData := []string{}
	mid := l / 2
	for n := 0; n < 5; n++ {
		// 洗牌
		for i := 0; i < mid; i += 2 {
			tmp := data[i]
			data[i] = data[mid+i]
			data[mid+i] = tmp
		}

		// 随机切牌
		for j := 0; j < 5; j++ {
			start := getRand()%(l-1) + 1
			numCards := getRand()%(l/2) + 1
			if start+numCards > l {
				numCards = l - start
			}

			tmpData = tmpData[:0]
			tmpData = append(tmpData, data[start:start+numCards]...)
			tmpData = append(tmpData, data[:start]...)
			tmpData = append(tmpData, data[start+numCards:]...)
		}
	}
	return tmpData
}
