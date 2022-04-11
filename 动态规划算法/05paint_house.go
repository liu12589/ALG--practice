package main

import "fmt"

/*
序列+状态
1、确定状态，最后一步如何花费最小呢，之前花费最小额度 + 当前剩余可选的最小额度。如何确定上一个颜色呢？
   当最后一种颜色为红色时，花费最小值为 dp[i][0] = min(dp[i - 1][1], dp[i - 1][2]) + costs[i][0]
   当最后一种颜色为蓝色时，花费最小值为 dp[i][1] = min(dp[i - 1][0], dp[i - 1][2]) + costs[i][1]
   当最后一种颜色为绿色时，花费最小值为 dp[i][2] = min(dp[i - 1][0], dp[i - 1][1]) + costs[i][2]
   最后比较dp[i][0]、dp[i][1]、dp[i][2]中的最小值即可
2、状态方程以列出（状态方程即前一步到这一步如何计算的）
3、初始条件和边界：注意dp[i][0]为前 i+1 个房子最小花费总金额。dp[0][0]前1个房子总金额。costs[i][0] 是第i+1个房子金额
4、顺序计算
*/

func main() {
	costs := [][]int{{3, 5, 3}, {6, 17, 6}, {7, 13, 18}, {9, 10, 18}}

	length := len(costs)
	var dp = make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, 3)
	}

	dp[0][0] = costs[0][0]
	dp[0][1] = costs[0][1]
	dp[0][2] = costs[0][2]

	for i := 1; i < length; i++ {
		dp[i][0] = min2(dp[i-1][1], dp[i-1][2]) + costs[i][0]
		dp[i][1] = min2(dp[i-1][0], dp[i-1][2]) + costs[i][1]
		dp[i][2] = min2(dp[i-1][0], dp[i-1][1]) + costs[i][2]
	}
	fmt.Println(min2(min2(dp[length-1][0], dp[length-1][1]), dp[length-1][2]))
}
func min2(a, b int) int {
	if a > b {
		return b
	}
	return a
}
