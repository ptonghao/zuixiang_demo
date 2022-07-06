/*
 * @Author: Jimpu
 * @Description: 约瑟夫 demo
 */

package josephus

// n个人,每数到m就出圈
func Josephus(n int, m int) int {
	if n == 1 {
		return 0
	} else {
		return (Josephus(n-1, m) + m) % n
	}
}
