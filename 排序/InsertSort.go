package main

import "fmt"

// 算法思想：每次将当前值与之前排好序的队列比较
// 降序
func insertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j] >= arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}

func main() {
	arr := []int{2, 4, 5, 1, 6}
	insertSort(arr)
	fmt.Println(arr)
}
