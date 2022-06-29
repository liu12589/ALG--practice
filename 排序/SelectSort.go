package main

import "fmt"

// SelectSort 每次遍历剩余位置数据，记录最大值放在当前位置
// 降序
func SelectSort(arr []int) {
	var curMin int
	for i := len(arr) - 1; i > 0; i-- {
		curMin = i
		for j := 0; j <= i; j++ {
			if arr[j] < arr[curMin] {
				curMin = j
			}
		}
		arr[i], arr[curMin] = arr[curMin], arr[i]
	}
}

func main() {

	arr := []int{1, 5, 4, 2}
	SelectSort(arr)
	fmt.Println(arr)

}
