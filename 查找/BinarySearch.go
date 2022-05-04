package main

import "fmt"

// 前提是有序的数组
// 双指针：start和end指向第一个节点和最后一个节点
// 1、等于中间节点值返回
// 2、大于中间节点值 i。start = i+ 1
// 3、小于中间节点值 i。end = i - 1
func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	n := binarySearch(nums, 12)
	fmt.Println(n)
}

func binarySearch(arr []int, n int) int {
	start, end := 0, len(arr)-1
	for start <= end {
		middle := (end-start)/2 + start
		if arr[middle] == n {
			return middle
		} else if arr[middle] < n {
			start = middle + 1
		} else {
			end = middle - 1
		}
	}
	return -1
}
