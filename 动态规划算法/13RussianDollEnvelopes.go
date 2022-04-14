package main

import (
	"fmt"
	"sort"
)

/*
给定n个信封的长度和宽度
如果一个信封的长和宽能够小于另一个信封的长和宽，则这个信封能够放入另一个信封
问最多能够嵌套多少个信封
——————————————————————————————————————————————————————————————
分析和确定状态：记录n-1个信封的最大最小值。
状态转移方程：sum = sum + if arr[i][0] > f[0][0] && arr[i][1] > f[0][1] || arr[i][0] < f[1][0] && arr[i][1] < f[1][1]
           f[0][0] = min{ f[0][0],arr[i][0] }
		   f[0][1] = min{ f[0][1],arr[i][1] }
		   f[1][0] = max{ f[1][0],arr[i][0] }
		   f[1][1] = max{ f[1][1],arr[i][1] }
初始值、边界条件：sum = 1； // 第一封信永远可以放进去
		   f[0][0] = arr[0][0]
		   f[0][1] = arr[0][1]
		   f[1][0] = arr[0][0]
		   f[1][1] = arr[0][1]
顺序计算
以上思路是错误的。没有对比 i 和 i-1 只能放入一个时放入哪个信封更合适。
做以下更改，以第i个信封为初始值。仍然不行。错误思路
———————————————————————————————————————————————————————————————
分析和确定状态：先按照长度排序
要求以第i个为最外层信封时最多的嵌套数，那么就需要知道以第j（j < i)个为最外层信封时的最多嵌套数
先按照长度排序 方便寻找
状态转移方程：f[i] = max{1,f[j] + 1|j 能否放入 i 中}
计算顺序：这个首先要排序，然后顺序计算
以上的解题时间复杂度为O(n^2)
———————————————————————————————————————————————————————————————
先按照长度排序后，将相同长度的逆序，相当于求最长递增子序列
*/
func main() {
	var arr = [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}
	fmt.Println(maxEnvelopes(arr))
}

func maxEnvelopes(arr [][]int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}
	sort.Slice(arr, func(i, j int) bool { return arr[i][0] < arr[j][0] })
	var length = len(arr)
	var f = make([]int, length)
	res := 0
	for i := 0; i < length; i++ {
		f[i] = 1
		for j := 0; j < i; j++ {
			if arr[j][0] < arr[i][0] && arr[j][1] < arr[i][1] {
				f[i] = max5(f[i], f[j]+1)
			}
		}
		res = max5(res, f[i])
	}
	return res
}

func max5(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func errFunc(arr [][]int) {
	var length = len(arr)
	var f = make([][]int, 2)
	for i := 0; i < 2; i++ {
		f[i] = make([]int, 2)
	}

	sum := 1
	res := 1
	for j := 0; j < length; j++ {
		f[0][0] = arr[j][0]
		f[0][1] = arr[j][1]
		f[1][0] = arr[j][0]
		f[1][1] = arr[j][1]
		for i := 1; i < length; i++ {
			if arr[i][0] < f[0][0] && arr[i][1] < f[0][1] {
				f[0][0] = arr[i][0]
				f[0][1] = arr[i][1]
				sum++
			} else if arr[i][0] > f[1][0] && arr[i][1] > f[1][1] {
				f[1][0] = arr[i][0]
				f[1][1] = arr[i][1]
				sum++
			}
		}
		if sum > res {
			res = sum
		}
		sum = 0
	}

	fmt.Println(res)
}
