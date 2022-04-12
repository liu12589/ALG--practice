package main

import (
	"fmt"
	"math"
)

/*
问题：有 N 栋房子有 M 种颜色，每两个房子之间颜色不能相同，二维数组表示每个房子涂哪种颜色的花费是多少。
——————————————————————————————————————————————————————————————————————————————————————
分析问题、确定状态：最后一个房子 + 之前花费最少的。依次假设最后一个房子颜色 + 之前花费最少的。然后遍历求解最小值
状态转移方程：f[N][x] = min(i!=x){f[N-1][i]} + arr[N-1][x]
初始值、边界条件：f[0][0]=0、f[1][x] = arr[0][x]
顺序计算
*/
func main() {
	N := 4 // 多少个房子
	M := 3 // 多少种颜色
	arr := [][]int{{3, 5, 3}, {6, 17, 6}, {7, 13, 18}, {9, 10, 18}}

	var f = make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		f[i] = make([]int, M)
	}

	for i := 1; i <= N; i++ {
		for j := 0; j < M; j++ {
			f[i][j] = min3(j, M, f[i-1]) + arr[i-1][j]
		}
	}

	// 遍历获取最小值
	var result = math.MaxInt64
	for i := 0; i < M; i++ {
		if f[N][i] < result {
			result = f[N][i]
		}
	}
	fmt.Println(result)
}

// 功能是选择f[N-1][i] 中除了 f[N-1][j] 的最小值
func min3(j, M int, f []int) int {
	var result = math.MaxInt64
	for i := 0; i < M; i++ {
		if i == j {
			continue
		}
		if f[i] < result {
			result = f[i]
		}
	}
	return result
}
