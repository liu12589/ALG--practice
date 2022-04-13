package main

import (
	"fmt"
)

/*
有一排N栋房子（0 - N-1），房子i里有A[i]个金币
一个窃贼想选择一些房子偷金币
但是不能偷任何挨着的两家邻居，否则会被警察抓住，问最多能偷多少金币。
————————————————————————————————————————————————————————————————————————————————————
分析（确定状态）：分析最后一次，前一家是否被偷如果被偷了，这个状态是不能确定的，类比颜色房子
			   设sum[i][0]是第i个房子没被偷的被偷最大金额 sum[i][0] = max{sum[i-1][0],sum[i-1][1]}
               sum[i][1] = sum[i-1][0] + arr[i-1]
状态转移方程：如上
初始值、边界条件：sum[0][0] = 0
顺序计算
———————————————————————————————————————————————————————————————————————————————————
优化1：sum[i][0]是第 i 个房子没被偷的最大金额，也就是相当于 i 房子不存在。可以做这样的变形
sum[i] = max { sum[i-1], sum[i-2] + arr[i-1] } 即取第 i 个房子没被偷和被偷的最大值。
优化2：只用三个变量存储当前now，前一个房子pre，前前一个房子prepre。

========================================================================================
题目变形：如果是一圈房子其他条件不变。
只判断最后一个节点：如果最后一个房子没被偷，那么之前一样
				 如果最后一个房子被偷，第一个房子没被偷，和之前一样
                 如果最后一个房子被偷，第一个房子也被偷，不能出现这种情况
如何避免最后一个房子和第一个房子同时被偷呢，分两次计算第一次从[1,n-1]、第二次从[2,n-1]
*/
func main() {
	arr := []int{15, 5, 13, 8, 9, 10}
	var length = len(arr)
	HouseRobber1(length, arr)
	HouseRobber2(length, arr)
	HouseRobber3(length, arr)
	HouseRobber4(length, arr)
}

func HouseRobber1(length int, arr []int) {
	var sum = make([][]int, length+1)
	for i := 0; i < length+1; i++ {
		sum[i] = make([]int, 2)
	}

	sum[0][0] = 0
	for i := 1; i <= length; i++ {
		sum[i][0] = max(sum[i-1][0], sum[i-1][1])
		sum[i][1] = sum[i-1][0] + arr[i-1]
	}
	if sum[length][0] > sum[length][1] {
		fmt.Println(sum[length][0])
	}
	fmt.Println(sum[length][1])
}

// HouseRobber2 优化后的方案
func HouseRobber2(length int, arr []int) {
	var sum = make([]int, length+1)

	sum[0] = 0
	sum[1] = arr[0]
	for i := 2; i <= length; i++ {
		sum[i] = max(sum[i-1], sum[i-2]+arr[i-1])
	}
	fmt.Println(sum[length])
}

// HouseRobber3 优化后的方案
func HouseRobber3(length int, arr []int) {
	var now = 0
	var pre = arr[0]
	var prepre int
	for i := 2; i <= length; i++ {
		now = max(pre, prepre+arr[i-1])
		prepre = pre
		pre = now
	}
	fmt.Println(now)
}

func HouseRobber4(length int, arr []int) {
	var sum = make([]int, length+1)
	var maxNum int

	sum[0] = 0
	sum[1] = arr[0]
	for i := 2; i <= length-1; i++ {
		sum[i] = max(sum[i-1], sum[i-2]+arr[i-1])
	}
	maxNum = sum[length-1]

	sum[1] = 0
	sum[2] = arr[1]
	for i := 3; i <= length; i++ {
		sum[i] = max(sum[i-1], sum[i-2]+arr[i-1])
	}
	if maxNum < sum[length] {
		maxNum = sum[length]
	}
	fmt.Println(maxNum)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
