package main

import "fmt"

// 无重复字符的最长子串 https://leetcode.cn/problems/longest-substring-without-repeating-characters/
// 思路；滑动窗口的意义在于，一、右指针向右移动找到第一个重复的值 二、左指针指向第一个重复值+1
func main() {
	s := "pwwkew"
	number := lengthOfLongestSubstring(s)
	fmt.Println(number)
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var left, right = 0, 0
	var max1 = 1
	for i := 1; i < len(s); i++ {
		for left <= right && s[i] != s[right] {
			length := i - right + 1
			if length > max1 {
				max1 = length
			}
			right--
		}
		left = right + 1
		right = i
	}
	return max1
}
