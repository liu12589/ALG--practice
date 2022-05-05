package main

import "fmt"

// 关键是在二分查找的基础上，找到出口的判定条件
func main() {
	nums := []int{4, 3, 2, 1}
	number := findMin(nums)
	fmt.Println(number)
}

func findMin(nums []int) int {
	start, end := 0, len(nums)-1
	result := nums[0]
	for start < end {
		middle := (end-start)/2 + start
		if nums[middle] >= result {
			if middle+1 < len(nums) && nums[middle+1] < nums[middle] {
				return nums[middle+1]
			}
			start = middle + 1
		} else {
			if middle-1 > 0 && nums[middle] < nums[middle-1] {
				return nums[middle]
			}
			end = middle - 1
			result = nums[middle]
		}
	}
	return result
}
