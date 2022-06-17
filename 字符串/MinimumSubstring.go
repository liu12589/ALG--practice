package main

import "fmt"

func main() {
	//s := "ADOBECODEBANC"
	//t := "ABC"
	s := "a"
	t := "aa"
	fmt.Println(minWindow(s, t))
}

// 退出条件：length = len(t)
// end 指针依次遍历 s，每遍历一个放入窗口中，等符合条件后，缩小窗口
// tMap[s[start]]、needLength 比较绕的是用两个变量巧妙的作为判定条件
func minWindow(s string, t string) string {
	var tMap = make(map[byte]int, len(s))
	for i := range t {
		tMap[t[i]]++
	}
	needLength := len(t)
	var resultStart, resultEnd = -1, len(s) + 1
	var start, end = 0, 0
	for end = range s {
		if tMap[s[end]] > 0 {
			needLength--
		}
		tMap[s[end]]--

		if needLength == 0 && start <= end {
			// 缩小范围
			for {
				if tMap[s[start]] == 0 {
					break
				}
				tMap[s[start]]++
				start++
			}
			if end-start < resultEnd-resultStart {
				resultStart, resultEnd = start, end
			}
			tMap[s[start]]++
			needLength++
			start++
		}
	}
	if resultStart == -1 {
		return " "
	}
	return s[resultStart : resultEnd+1]
}
