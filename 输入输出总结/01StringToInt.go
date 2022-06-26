package main

import "fmt"

func main() {
	s := 2342123
	// 用一个数组存放数据
	var arr []int
	for s > 0 {
		number := s % 10
		s /= 10
		arr = append(arr, number)
	}
	// 每个数依次相乘
	var result = 1
	for _, k := range arr {
		result *= k
	}
	fmt.Println(result)
}
