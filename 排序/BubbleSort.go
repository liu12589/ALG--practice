package main

import "fmt"

// BubbleSort 冒泡排序
// 将j位置数与j+1位置比较，j大则交换，小不交换，j++。
// 注意一次循环完成后，判定是有序的直接退出
// 降序
func BubbleSort(arr []int) {
	for i := len(arr) - 1; i > 0; i-- {
		isSort := true
		for j := 0; j < i && arr[j] < arr[j+1]; j++ {
			arr[j], arr[j+1] = arr[j+1], arr[j]
			isSort = false
		}
		if isSort {
			break
		}
	}
}

func main() {
	arr := []int{2, 4, 6, 3, 1}
	BubbleSort(arr)
	fmt.Println(arr)
}
