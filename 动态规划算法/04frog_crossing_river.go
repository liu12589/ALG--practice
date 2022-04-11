package main

import "fmt"

/*
j 为第 n 个石头，i 为要判断的能否跳到 j 的石头
设 k = stones[j] - stones[i]
1、要想判断能否跳到第 n 个，需要判断 k > j + 1 范围内 i 能否跳到 n。还需要判断能否到达 i。
2、f[i][k] 表示第 i 个石头是否能否跳到距离为k的石头上。
3、状态转移方程为 f[i][k] = OR(0<= i < j)(dp[i][k] = dp[j][k-1] || dp[j][k] || dp[j][k+1])
*/
func main() {
	stones := []int{0, 1, 3, 5, 6, 8, 12, 17}
	n := len(stones)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	dp[0][0] = true
	for i := 1; i < n; i++ {
		if stones[i]-stones[i-1] > i {
			fmt.Println(0)
		}
	}
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			k := stones[i] - stones[j]
			if k > j+1 {
				break
			}
			dp[i][k] = dp[j][k-1] || dp[j][k] || dp[j][k+1]
			if i == n-1 && dp[i][k] {
				fmt.Println(1)
			}
		}
	}
	fmt.Println(0)
}
