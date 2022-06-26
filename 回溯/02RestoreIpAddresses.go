package main

import (
	"fmt"
	"math"
)

// 简化，将所有情况输出
func main() {
	s := "25525511135"
	fmt.Println(restoreIpAddresses(s))
}

func restoreIpAddresses(s string) [][]string {
	var arr [][]string
	var IpAddresses func(curS string, temp []string)
	IpAddresses = func(curS string, temp []string) {
		if len(temp) == 3 {
			temp = append(temp, curS)
			newarr := make([]string, len(temp))
			copy(newarr, temp)
			arr = append(arr, newarr)
			temp = []string{}
			return
		}
		for i := 0; i < 3; i++ {
			// 排除这一个长度后
			curSLength := len(curS) - i - 1
			tempLength := 3 - len(temp)
			// 剩余字符串长度 / 数组能够容纳量 向上取整
			cur := int(math.Ceil(float64(curSLength) / float64(tempLength)))
			if curSLength < tempLength || cur > 3 {
				continue
			}
			str := curS[:i+1]
			temp = append(temp, curS[:i+1])
			curS = curS[i+1:]
			IpAddresses(curS, temp)
			curS += str
		}
	}
	IpAddresses(s, []string{})
	return arr
}
