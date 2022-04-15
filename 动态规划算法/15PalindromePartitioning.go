package main

import (
	"fmt"
	"math/big"
)

/*
给定一个字符串 s，请将 s 分割成一些子串，使每个子串都是回文串。
返回符合要求的 最少分割次数 。
——————————————————————————————————————————————————————
f[i] = min{f[j]&&arr[j][i-1]==true}  + 1
这里的arr[j][i-1]指下标从j到i-1的子串是否是回文数
初始值：f[0]=0 f[1] = 1。边界：从0到n-1。第二次遍历i 从0到i
顺序计算
*/
func main() {
	s := "aab"
	length := len(s)
	var f = make([]int, length+1)
	arr := allIsPalindrome(s)

	f[0] = 0
	for i := 1; i <= length; i++ {
		f[i] = big.MaxExp
		for j := 0; j < i; j++ {
			if arr[j][i-1] == true && f[j] < f[i] {
				f[i] = f[j] + 1
			}
		}
	}
	fmt.Println(f[length] - 1)
}

// 可以判断回文串，但是判断出所有的回文串效率太低
func isPalindrome(str string, start int, end int) bool {
	for start < end {
		if str[start] != str[end] {
			return false
		}
		start++
		end--
	}
	return true
}

// 求所有的回文子串
func allIsPalindrome(str string) [][]bool {
	length := len(str)
	var isPalin = make([][]bool, length)
	for i := 0; i < length; i++ {
		isPalin[i] = make([]bool, length)
	}
	for i := 0; i < length; i++ {
		for h := 0; h < length; h++ {
			isPalin[i][h] = false
		}
	}
	var i, j, t int
	for t = 0; t < length; t++ {
		i = t
		j = t
		for i >= 0 && j < length && str[i] == str[j] {
			isPalin[i][j] = true
			i--
			j++
		}
		i = t
		j = t + 1
		for i >= 0 && j < length && str[i] == str[j] {
			isPalin[i][j] = true
			i--
			j++
		}
	}
	return isPalin
}
