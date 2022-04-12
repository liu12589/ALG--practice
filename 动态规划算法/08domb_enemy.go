package main

import "fmt"

/*
题目：有 M*N 个网格，每个格子可能是空的，可能有一个敌人，也可能有一堵墙、也可能是空的
	 只能在空地放炸弹，炸弹会炸死同行同列的敌人，但是不能炸透墙。
     求最多能够炸死多少个敌人
——————————————————————————————————————————————————————————————————————
题目分析（确定状态）:先计算该格上方能够炸死多少敌人，将值累加到结果中，重复计算左，右，下。
这里要开两个二维数组，一个f[m][n]是为了记录上下左右能够炸死多少人，一个sum[m][n]是保存炸死总人数
状态方程：up f[m][n] ={ 0(if f[m-1][n] = w) V f[m-1][n] + 1(if f[m-1][n] = E) V f[m-1][n](if f[m-1][n] = "0")
		sum[m][n] = up f[m][n] + down f[m][n] + left f[m][n] + right f[m][n]
初始值、边界 up 且n = 0时 f[m][n] = 0
		   down且 n = N -1 时 f[m][n] = 0
           right且m = M -1 时 f[m][n] = 0
           left 且m = 0 时 f[m][n] = 0
计算顺序：上 从f[0][0]开始，先向左遍历，再向下遍历
		下 从f[0][n-1]开始，先向左遍历，再向上遍历
		左 从f[0][0]开始，先向左遍历，再向下遍历
		右 从f[m-1][0]开始，先向右遍历，再向下遍历
==============================================================================
一定要注意细节！！！！！！！！！
*/
func main() {

	var arr = [][]byte{{'0', 'E', '0', '0'},
		{'E', '0', 'W', 'E'},
		{'0', 'E', '0', '0'}}
	// m 代表有多少列，n 代表有多少行
	var m = len(arr[0])
	var n = len(arr)

	var f = make([][]int, n)
	var sum = make([][]int, n)
	for i := 0; i < n; i++ {
		sum[i] = make([]int, m)
		f[i] = make([]int, m)
	}

	// 计算向上炸死的人数
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if arr[i][j] == 'W' {
				f[i][j] = 0
			} else {
				f[i][j] = 0
				if arr[i][j] == 'E' {
					f[i][j] = 1
				}
				if i > 0 {
					f[i][j] += f[i-1][j]
				}
			}
			sum[i][j] += f[i][j]
		}

	}

	// 计算向下炸死的人数
	for i := n - 1; i >= 0; i-- {
		for j := 0; j < m; j++ {
			if arr[i][j] == 'W' {
				f[i][j] = 0
			} else {
				f[i][j] = 0
				if arr[i][j] == 'E' {
					f[i][j] = 1
				}
				if i < n-1 {
					f[i][j] += f[i+1][j]
				}
			}
			sum[i][j] += f[i][j]
		}
	}

	// 计算左侧炸死的人数
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if arr[i][j] == 'W' {
				f[i][j] = 0
			} else {
				f[i][j] = 0
				if arr[i][j] == 'E' {
					f[i][j] = 1
				}
				if j > 0 {
					f[i][j] += f[i][j-1]
				}
			}
			sum[i][j] += f[i][j]
		}
	}

	var result = 0
	// 计算右侧炸死的人数,
	for i := 0; i < n; i++ {
		for j := m - 1; j >= 0; j-- {
			if arr[i][j] == 'W' {
				f[i][j] = 0
			} else {
				f[i][j] = 0
				if arr[i][j] == 'E' {
					f[i][j] = 1
				}
				if j < m-1 {
					f[i][j] += f[i][j+1]
				}
			}
			sum[i][j] += f[i][j]
			if sum[i][j] > result && arr[i][j] == '0' {
				result = sum[i][j]
			}
		}
	}
	fmt.Println(result)
}
