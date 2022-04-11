package main

import "fmt"

/*
给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
输出：7
解释：因为路径 1→3→1→1→1 的总和最小。
-----------------------------------------------------
1、确定状态
我们可以把 grid 存放到达每一步数字总和最小。最后一步可有两个选择：grid[column-1][row]、grid[column][row-1]
化解成子问题，可以理解为：到达最后一格总和最小即选择前一格最小值+本格的值。即：grid[column][row] = min(grid[column][row-1], grid[column-1][row]) + grid[column][row]

2、确定转移方程 grid[i][j] = min{grid[i][j-1], grid[i-1][j]} + grid[i][j]

3、确定初始条件和边界
- 无初始条件
- 边界: grid[0][0] 无需计算，遇到时直接跳过。
- 当 column = 0 时只能是本行前一格到达。grid[i][j] = grid[i][j-1] + grid[i][j]
- 当 row = 0 时只能是本列上一格到达。grid[i][j] = grid[i-1][j] + grid[i][j]
- 确定计算顺序：先计算行计算，再列计算

*/
func main() {

	const m, n = 3, 3
	// 注意初始化数组时，其规定长度必须是定长，否则报错
	var f = [m][n]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}

	f[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if j == 0 {
				f[i][j] = f[i-1][j] + f[i][j]
			} else if i == 0 {
				f[i][j] = f[i][j-1] + f[i][j]
			} else if i != 0 {
				f[i][j] = min1(f[i-1][j], f[i][j-1]) + f[i][j]
			}
		}
	}
	fmt.Println(f[m-1][n-1])
}

func min1(a, b int) int {
	if a > b {
		return b
	}
	return a
}
