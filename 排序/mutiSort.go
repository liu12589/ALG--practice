package main

import (
	"fmt"
	"math"
)

// 将一个二维升序的数组，合并为升序的一维数组

func main() {
	arr := [][]int{{}, {2, 3, 7}, {4, 5, 7}, {1}}
	var result []int
	var min int
	for len(arr) > 0 {
		curMin := math.MaxInt32
		for k := 0; k < len(arr); {
			if len(arr[k]) == 0 {
				arr = append(arr[:k], arr[k+1:]...)
				continue
			}
			if len(arr[k]) > 0 && arr[k][0] < curMin {
				curMin = arr[k][0]
				min = k
			}
			k++
		}
		result = append(result, curMin)
		arr[min] = arr[min][1:]
		if len(arr[min]) == 0 {
			arr = append(arr[:min], arr[min+1:]...)
		}
	}
	fmt.Println(result)
}
