package main

import "fmt"

// f[n]表示在n个位置时爱丽丝先手能否胜利，知道其能胜利，就要知道 f[n-j](n % j == 0)的状态，如果f[n-j] == false，那么 f[n]就为true
// f[n] = 存在 f[n-j]=false，其中 n % j == 0 且 0 < j < n

func main() {
	relation := [][]int{{0, 2}, {2, 1}, {3, 4}, {2, 3}, {1, 4}, {2, 0}, {0, 4}}

	for _, r := range relation {
		fmt.Println(r)
	}
}

func longestPalindrome(s string) string {
	length := len(s)
	var f = make([][]bool, length)
	for i := range f {
		f[i] = make([]bool, length)
	}

	var m, n int
	for i := 0; i < length; i++ {
		m = i
		n = i
		for m >= 0 && n < length && s[m] == s[n] {
			f[m][n] = true
			m--
			n++
		}

		m = i
		n = i + 1
		for m >= 0 && n < length && s[m] == s[n] {
			f[m][n] = true
			m++
			n++
		}
	}
	var result = -1
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			if f[i][j] && j-i > result {
				m = i
				n = j
				result = j - i
			}
		}
	}
	fmt.Println(f)
	fmt.Println(m, n)
	return string(s[m : n+1])
}

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	//保存前n-1个数之和 + 第n个数的最大值
	var f = make([]int, len(nums))
	//保存全局最大值
	var result int

	f[0] = nums[0]
	result = f[0]
	for i := 1; i < len(nums); i++ {
		f[i] = max(f[i-1]+nums[i], nums[i])
		if f[i] > result {
			result = f[i]
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
