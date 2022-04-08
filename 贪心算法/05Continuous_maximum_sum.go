package main

import "fmt"

func main() {
	var num, number int
	var arr []int
	fmt.Scan(&num)
	for i := 0; i < num; i++ {
		fmt.Scan(&number)
		arr = append(arr, number)
	}
	maximum(arr)

}

func maximum(arr []int) {
	max := arr[0]
	maxSum, result := arr[0], arr[0]
	for i := 0; i < len(arr); i++ {
		max = maxSum + arr[i]
		maxSum = max1(arr[i], max)
		result = max1(result, maxSum)
	}
	fmt.Println(result)
}

func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}
