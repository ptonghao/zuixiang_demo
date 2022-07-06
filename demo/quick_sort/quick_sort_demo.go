/*
 * @Author: Jimpu
 * @Description: 实现一个函数，输入为任意长度的int数组，返回元素最大差值，例如输入arr=[5,8,10,1,3]，返回9
 */

package quick_sort

// quickSort 快速排序
func QuickSort(arr []int, left, right int) {
	if left < right {
		pivot := arr[left]
		j := left
		for i := left; i < right; i++ {
			if arr[i] < pivot {
				j++
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
		arr[left], arr[j] = arr[j], arr[left]
		QuickSort(arr, left, j)
		QuickSort(arr, j+1, right)
	}
}
