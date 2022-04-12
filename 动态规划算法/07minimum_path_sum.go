package main

import "fmt"

/*
问题1：求最长连续递增子序列
输入： 4 1 2 3
输出： 3 （因为 1 2 3）
-------------------------------------------------------------------
分析（确定状态）：设f[n]表示包含第n个字符的最长单调长度、然后取f[n]数组中最大值
确定状态转移方程：f[n] = max {1, f[n-1]+1 |arr[i] > arr[i-1]}
确定初始值、边界：f[0] = 1
顺序计算
====================================================================
问题2：求最长连续单调子序列
输入： 4 1 2 3 2 1 0
输出： 4 （因为 3210）
不要考虑的太复杂了，递增的求完了，保留最大值。再将字符串逆序，再求一遍保留最大值。
时间复杂度 O(n)、空间复杂度O(n)
====================================================================
问题3：如何优化成空间复杂度O(1)
分析：1、设一个全局变量 max，每次保留最大值。不用最后取f[n]数组中最大值
	 2、那么只需要保留包含前一个字符的最大长度子串
*/
var maxInt = 0

func main() {
	arr := "4321"
	var arrInt = make([]int, len(arr))
	for i, v := range []byte(arr) {
		v -= '0'
		arrInt[i] = int(v)
	}
	num1 := minimumPath1(arrInt)
	reverse(arrInt)
	num2 := minimumPath1(arrInt)
	if num1 > num2 {
		fmt.Println(num1)
	} else {
		fmt.Println(num2)
	}
}

func minimumPath1(arrInt []int) int {
	length := len(arrInt)
	f := 1
	for i := 1; i < length; i++ {
		f = 1
		if arrInt[i] > arrInt[i-1] {
			f = f + 1
		}
		if maxInt < f {
			maxInt = f
		}
	}
	return maxInt
}

func reverse(arr []int) {
	i := 0
	j := len(arr) - 1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}
