package main

import "fmt"

// 使用滑动窗口
func main() {
	s1 := "hello"
	s2 := "ooolleoooleh"
	fmt.Println(checkInclusion1(s1, s2))
}

// 固定长度滑动窗口
// 简化为固定一个 s1 长度的滑动窗口，每次对比是否符合条件
// 这类问题的关键是如何巧妙的判断相符合的条件啊

func checkInclusion(s1 string, s2 string) bool {
	lengthS1 := len(s1)
	lengthS2 := len(s2)
	if lengthS1 > lengthS2 {
		return false
	}
	arr := [26]int{}
	for _, v := range s1 {
		arr[v-'a']--
	}
	var start, end = 0, 0
	for end = range s2 {
		word := s2[end] - 'a'
		arr[word]++
		for arr[word] > 0 {
			arr[s2[start]-'a']--
			start++
		}
		if end-start+1 == lengthS1 {
			return true
		}
	}
	return false
}

func checkInclusion1(s1 string, s2 string) bool {
	lengthS1 := len(s1)
	lengthS2 := len(s2)
	needLength := lengthS1 // 记录固定窗口的大小
	if lengthS1 > lengthS2 {
		return false
	}
	var s1Map = make(map[byte]int, lengthS2)
	for i := range s1 {
		s1Map[s1[i]]++
	}
	var start = 0
	// 使用固定窗口，以start的遍历为依据
	for end := range s2 {
		if s1Map[s2[end]] > 0 {
			needLength--
		}
		s1Map[s2[end]]--
		// 长度一致，先判断是否符合条件，符合条件直接返回。不符合条件start向右移动一位
		if end-start+1 == lengthS1 {
			if needLength == 0 {
				return true
			}
			if s1Map[s2[start]] >= 0 {
				needLength++
			}
			s1Map[s2[start]]++
			start++
		}
	}
	return false
}
