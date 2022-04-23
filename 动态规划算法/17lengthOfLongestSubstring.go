package main

import "fmt"

// 这个题虽然没有用动态规划常规解法，但是用了其思想。保留包含i-1位时符合条件的最长子串(这里用双指针，而不是数组保存）。
// 遍历到i时，对比i-1子串，求包含i位时符合条件的最长子串，然后记录该子串。最后返回最大值。
// 双指针问题：start 指向符合条件的包含第i-1个子串的初始位置；end 指向符合条件的包含第i-1个子串的结束位置
// 遍历求以第i个元素结尾的符合条件的子串长度
// 遍历第i-1个子串，从start到end。从i-1个结尾向前当 end < start || s[end] == s[i] 时结束。记录最大值
func main() {
	s := "pwwkew"
	result := lengthOfLongestSubstring(s)
	fmt.Println(result)
}

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	start, end := 0, 0
	max1 := 1
	for i := 1; i < len(s); i++ {
		for end >= start && s[end] != s[i] {
			length := i - end + 1
			if length > max1 {
				max1 = length
			}
			end--
		}
		start = end + 1
		end = i
	}
	return max1
}
