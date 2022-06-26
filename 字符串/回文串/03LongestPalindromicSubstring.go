package main

import "fmt"

// 双指针
func main() {
	arr := "babad"
	a := longestPalindrome(arr)
	fmt.Println(a)
}

// 初始化两个指针，两个保存最终结果
func longestPalindrome(s string) string {
	max := 0
	length := len(s)
	left, right := 0, 0
	left1, right1 := 0, 0
	for i := 1; i < length; i++ {
		left, right = i, i
		for left >= 0 && right < length && s[left] == s[right] {
			gap := right - left
			if max < gap {
				max = gap
				left1 = left
				right1 = right
			}
			left--
			right++
		}
		left, right = i, i+1
		for left >= 0 && right < length && s[left] == s[right] {
			gap := right - left
			if max < gap {
				max = gap
				left1 = left
				right1 = right
			}
			left--
			right++
		}

	}
	return s[left1 : right1+1]
}
