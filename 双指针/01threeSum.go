package main

import (
	"fmt"
	"sort"
)

// 暴力求解
//排序，双指针
// 先对数组进行排序
// 遍历nums，如果遇到相同的数，跳过，如果nums[i] > 0直接退出循环
// left、right 指向i后边的i+1， length-1的位置，遍历剩余数组。

func main() {
	//nums := []int {-1, 0, 1, 2, -1, -4}
	nums := []int{-2, 0, 0, 2, 2}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) (result [][]int) {
	sort.Ints(nums)
	length := len(nums)
	if length <= 1 || nums[length-1] < 0 {
		return result
	}
	for i := 0; i < length; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := length - 1
		for left < right && nums[i] <= 0 {
			if left > i+1 && nums[left] == nums[left-1] {
				left++
				continue
			}
			if right < length-1 && nums[right] == nums[right+1] {
				right--
				continue
			}
			number := nums[i] + nums[left] + nums[right]
			if number == 0 {
				arr := []int{nums[i], nums[left], nums[right]}
				result = append(result, arr)
				left++
				right--
			} else if number < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}
