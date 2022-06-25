package main

import (
	"fmt"
	"math"
	"sort"
)

// 632. 最小区间 https://leetcode.cn/problems/smallest-range-covering-elements-from-k-lists/
func main() {
	nums := [][]int{{4, 10, 15, 24, 26}, {0, 9, 12, 20}, {5, 18, 22, 30}}
	//nums := [][]int{{1,2,3},{1,2,3},{1,2,3}}
	fmt.Println(smallestRange(nums))
}

// 分析：必须遍历完才知道最优解，最优解包含在数组中
// 关键：如何遍历，如何判断符合条件
func smallestRange(nums [][]int) []int {
	length := len(nums)
	var arr [][]int
	j := 0
	for index, k := range nums {
		for _, value := range k {
			var middleArr = make([]int, 2)
			middleArr[0] = value
			middleArr[1] = index
			arr = append(arr, middleArr)
			j++
		}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})
	var start = 0
	var needNum = make(map[int]int, length)
	var needLength = length
	var result = math.MaxInt32
	var resultArr = make([]int, 2)
	for end, k := range arr {
		if needNum[k[1]] == 0 {
			needLength--
		}
		needNum[k[1]]++
		for needLength == 0 {
			if arr[end][0]-arr[start][0] < result {
				result = arr[end][0] - arr[start][0]
				resultArr[0], resultArr[1] = arr[start][0], arr[end][0]
			}
			// 移动左指针
			startNumLocation := arr[start][1]
			needNum[startNumLocation]--
			if needNum[startNumLocation] == 0 {
				needLength++
			}
			start++
		}
	}
	return resultArr
}
