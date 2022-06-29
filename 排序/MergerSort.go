package main

import "fmt"

// 归并排序
// 核心思想：
// 使用递归拆分，直到为 1。
// 然后合并
func mergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	middle := length / 2
	left := arr[:middle]
	right := arr[middle:]
	return merge(mergeSort(left), mergeSort(right))
}

// 合并两个有序的数组
func merge(left, right []int) []int {
	var result []int
	var curLeft, curRight int
	for {
		if curLeft == len(left) {
			result = append(result, right[curRight:]...)
			break
		}
		if curRight == len(right) {
			result = append(result, left[curLeft:]...)
			break
		}
		if left[curLeft] < right[curRight] {
			result = append(result, left[curLeft])
			curLeft++
		} else {
			result = append(result, right[curRight])
			curRight++
		}
	}
	return result
}

func main() {
	arr := []int{1, 2, 44, 3, 5, 6}
	fmt.Println(mergeSort(arr))
}
