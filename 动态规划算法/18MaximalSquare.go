package main

import "fmt"

func main() {
	matrix := [][]byte{{'0', '0', '0', '1'}, {'1', '1', '0', '1'}, {'1', '1', '1', '1'}, {'0', '1', '1', '1'}, {'0', '1', '1', '1'}}
	number := maximalSquare(matrix)
	fmt.Println(number)
}

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])
	var f = make([][]int, m)
	for i := 0; i < m; i++ {
		var arr = make([]int, n)
		f[i] = arr
	}

	result := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 || f[i][j-1] == 0 || f[i-1][j] == 0 {
					f[i][j] = 1
				} else if f[i][j-1] != f[i-1][j] {
					f[i][j] = min11(f[i][j-1], f[i-1][j]) + 1
				} else if matrix[i-f[i-1][j]][j-f[i-1][j]] == '1' {
					f[i][j] = f[i-1][j] + 1
				} else {
					f[i][j] = f[i-1][j]
				}
			}
			if f[i][j] > result {
				result = f[i][j]
			}
		}
	}
	return result * result
}
func min11(a, b int) int {
	if a < b {
		return a
	}
	return b
}
