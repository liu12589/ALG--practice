package main

import "fmt"

// 思路：首先明白暴力搜索，
// 其实根二叉树深度遍历差不多
func main() {
	nums := []int{1, 2, 3}
	//permute(nums)
	fmt.Println(nums[0])
}
func permute(nums []int) [][]int {
	length := len(nums)
	var flag = make([]bool, length) //false 代表可以访问，true代表不可以访问
	var arr [][]int
	var trace func(curLength int, temp []int)
	trace = func(curLength int, temp []int) {
		if curLength == length {
			newarr := make([]int, len(temp))
			copy(newarr, temp)
			arr = append(arr, newarr)
			return
		}
		for j := 0; j < length; j++ {
			if flag[j] == true {
				continue
			} else {
				temp = append(temp, nums[j])
				flag[j] = true
				trace(curLength+1, temp)
				temp = temp[:len(temp)-1]
				flag[j] = false
			}
		}
	}
	trace(0, []int{})
	return arr
}
