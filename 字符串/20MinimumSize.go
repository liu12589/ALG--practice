package main

import "fmt"

func main() {
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen(target, nums))
}

// 如何判定连续
func minSubArrayLen(target int, nums []int) int {
	var start, end = 0, 0
	var sum int
	var result = len(nums) + 1
	for end < len(nums) {
		sum += nums[end]
		if sum >= target {
			for sum >= target {
				sum -= nums[start]
				start++
			}
			if result > end-start+2 {
				result = end - start + 2
			}
		}
		end++
	}
	if result == len(nums)+1 {
		return 0
	}
	return result
}
