package main

import "fmt"

// 思路：用数组存储匹配成功的字符串的初始位置，然后检查符合条件的位置
// 问题的关键在于怎么判定最终符合条件
func main() {
	arr := []string{"foo", "bar"}
	s := "barfoothefoobarman"
	//s := "wordgoodgoodgoodbestword"
	//arr := []string{"word","good","best","word"}
	fmt.Println(findSubstring(s, arr))

}

// 思想：左右指针，左指针不动，移动右指针，右指针不符合条件时，移动左指针
// 左右指针移动条件
// 左右指针如何移动
// 窗口退出条件，什么时候符合条件
func findSubstring(s string, words []string) (result []int) {
	var wordLength = len(words[0])
	var wordNumber = len(words)
	var wordResultLength = wordLength * wordNumber
	var start, end = 0, 0

	for i := 0; i < len(s)-wordResultLength+1; i++ {
		start = i
		end = i
		mapArr := deepCopy(words)
		for true {
			endWord := s[end : end+wordLength]
			if mapArr[endWord] > 0 {
				end += wordLength
				mapArr[endWord]--
			} else {
				break
			}
			if end-start == wordResultLength {
				result = append(result, start)
				break
			}
		}
	}
	return result
}

func deepCopy(words []string) map[string]int {
	var resultMap = make(map[string]int, len(words))
	for _, word := range words {
		resultMap[word]++
	}
	return resultMap
}
