package main

import (
	"fmt"
)

func main() {
	var arr []int
	var n int
	var str int
	fmt.Scanf("%d", &str)
	for str > 0 {
		arr = append(arr, str%10)
		str = str / 10
	}
	length := len(arr)
	for i := 0; i < length/2; i++ {
		temp := arr[length-1-i]
		arr[length-1-i] = arr[i]
		arr[i] = temp
	}
	fmt.Scanf("%d", &n)
	maxNumber(arr, n)
}

// 思路：每次取前low, low + n 个数，返回最大值 + 1及其位置
//下一次从 low，low + n 范围内找最大值
//如果low + 1 + n > length 则取length
//遍历 length - n 次
func maxNumber(arr []int, n int) {
	low, max := 0, 0
	var result int
	for i := 0; i < len(arr)-n; i++ {
		if low+n > len(arr)-1 {
			low, max = findMax(arr, low, len(arr)-1)
		} else {
			low, max = findMax(arr, low, low+n)
		}
		result = result*10 + max
	}
	fmt.Println(result)
}

func findMax(arr []int, low, high int) (int, int) {
	max := arr[low]
	for i := low; i <= high; i++ {
		if arr[i] > max {
			low = i
			max = arr[i]
		}
	}
	return low + 1, max
}
