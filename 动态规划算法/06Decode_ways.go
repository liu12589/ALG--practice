package main

import (
	"fmt"
)

/*
子问题：设 Sum[i] 是 i 个字符能够组成的类别。字符串为 S
状态转移方程为： Sum[i] = sum[i -1] ｜ S[i-1]对应一个字母 + sum[i-2] ｜ S[i-2]S[i-1]对应一个字母
边界条件和初始值：f[0]=1等于空串、i = 1 只看最后一个数字
*/

func main() {
	s := "1199"
	length := len(s)
	var arr = make([]int, length)
	for i, v := range []byte(s) {
		v -= '0'
		arr[i] = int(v)
	}

	var sum = make([]int, length+1)
	sum[0] = 1
	for i := 1; i <= length; i++ {
		sum[i] = 0
		if arr[i-1] >= 1 && arr[i-1] <= 9 {
			sum[i] += sum[i-1]
		}
		if i >= 2 {
			t := arr[i-2]*10 + arr[i-1]
			if t >= 10 && t <= 26 {
				sum[i] += sum[i-2]
			}
		}
	}
	fmt.Println(sum[length])
}
