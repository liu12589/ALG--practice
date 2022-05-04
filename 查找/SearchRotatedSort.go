package main

import "fmt"

// 依旧是二分查找，不过要判断向左查找还是向右查找
// arr[m] > arr[start] 不包含旋转,正常二分查找
// arr[m] > arr[start] 包含旋转，

func main() {
	arr := []int{4, 5, 6, 7, 0, 1, 2}
	n := binarySearch1(arr, 2)
	fmt.Println(n)
}

func binarySearch1(arr []int, n int) int {
	start, end := 0, len(arr)-1
	for start <= end {
		middle := (end-start)/2 + start
		if arr[middle] == n {
			return middle
		}
		if arr[middle] >= arr[0] {
			if arr[0] <= n && n < arr[middle] {
				end = middle - 1
			} else {
				start = middle + 1
			}
		} else {
			if arr[end] >= n && n > arr[middle] {
				start = middle + 1
			} else {
				end = middle - 1
			}
		}
	}
	return -1
}
