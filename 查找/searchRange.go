package main

// 范围查找，查找左右节点。最直观的解法就是利用二分查找，分别查找左右节点
// 查找时多加个调价本元素==target时，上一个元素小于target时，表明是left，查找right原理相同。
import "fmt"

func main() {
	nums := []int{1}
	fmt.Println(searchRange(nums, 1))
}

func searchRange(nums []int, target int) []int {
	start, end := 0, len(nums)-1
	left, right := -1, -1
	for start <= end {
		middle := (end-start)/2 + start
		//寻找右地址
		if nums[middle] == target {
			if middle+1 > len(nums)-1 || nums[middle+1] > target {
				right = middle
				break
			}
			start = middle + 1
		} else if nums[middle] < target {
			start = middle + 1
		} else {
			end = middle - 1
		}
	}
	start, end = 0, len(nums)-1
	for start <= end {
		middle := (end-start)/2 + start
		//寻找左地址
		if nums[middle] == target {
			if middle-1 < 0 || nums[middle-1] < target {
				left = middle
				break
			}
			end = middle - 1
		} else if nums[middle] < target {
			start = middle + 1
		} else {
			end = middle - 1
		}
	}
	result := []int{left, right}
	return result
}
