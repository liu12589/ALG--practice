package main

import (
	"fmt"
	"math"
)

// 测试map是引用拷贝

func main() {
	arr := []int{6, 3, 2, 1}
	push(arr, 4)

	var arrMap = make(map[int]int, 3)
	arrMap[1] = 10
	arrMap[0] = 20
	arrMap[3] = 30
	var curMax, curMin = math.MinInt32, math.MaxInt32
	for i, k := range arrMap {
		if k <= curMin {
			curMin = i
		}
		if k > curMax {
			curMax = i
		}
	}
	fmt.Println(curMax, curMin)

	// 测试二维数组
	var twoArr = make([][]int, 2)
	twoArr[0] = make([]int, 2)
	twoArr[1] = []int{1, 2}
	fmt.Println(twoArr[0][0])

	// 二维数组遍历
	for i, k := range twoArr {
		fmt.Println(i, k)
	}

	nums := []int{1, 2, 5, 6, 3, 4}
	fmt.Println(find(nums, 2))
}

func push(arr []int, number int) {
	for i, k := range arr {
		if number > k {
			arr = arr[:i]
			arr = append(arr, number)
			break
		}
	}
}

// 遍历一次找到数组中倒数第二个最小值
func find(nums []int, n int) (int, int) {
	var min = math.MaxInt32    // 记录最小值
	var middle = math.MaxInt32 // 记录次小值
	for _, k := range nums {
		// 小于最小的原，更新min，middle,小于middle更新middle
		if min > k {
			middle = min
			min = k
		} else if middle > k && k != min {
			middle = k
		}
	}
	return min, middle
}
