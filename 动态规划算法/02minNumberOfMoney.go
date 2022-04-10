package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
给定数组arr，arr中所有的值都为正整数且不重复。
每个值代表一种面值的货币，每种面值的货币可以使用任意张，再给定一个aim，代表要找的钱数，求组成aim的最少货币数。
示例1
输入：3 20
	 5 2 3
输出：4
说明：20=5*4
--------------------------------------------------------------------------------------------
题解：
1、确定状态。也就是确定最后一步如何选择是最优的 f(x) = min { f(x-5)、f(x-2)、f(x-3) } + 1
2、确定状态转移方程：f(x) = min { f(x-5)、f(x-2)、f(x-3) } + 1
3、确定边界、初始值（这一步很关键）这一题没有边界问题，初始值是哪些已知的变量。循环求解时这些值略过。
4、确定执行顺序（很关键）状态转移方程中变量一定是已知的。
5、写程序时要注意灵活多变，实际问题中求f(x) = min { f(x-5)、f(x-2)、f(x-3) }，而f(x-5)、f(x-2)是无法求解的，只能选择f(x-3)。
   因此，可以用最大值代替无法求解出的数，也可以用-1。但用了 -1 后打乱了之前逻辑，需要额外判断
*/

func main() {

	var aim, arrLength int
	fmt.Scan(&arrLength, &aim)
	var arr []int
	moneyStrs := read()
	for i := 0; i < len(moneyStrs); i++ {
		number, _ := strconv.Atoi(moneyStrs[i])
		arr = append(arr, number)
	}

	// 数组开多大的问题。0 ... n 开 n+1
	var f = make([]int, aim+1)
	f[0] = 0
	for i := 1; i <= aim; i++ {
		f[i] = -1
		// 求解组成 i 值最少个数
		for j := 0; j < arrLength; j++ {
			if i >= arr[j] && f[i-arr[j]] != -1 {
				f[i] = min(f[i], f[i-arr[j]]+1)
			}
		}
	}
	fmt.Println(f[aim])
}

func min(a, b int) int {
	if a == -1 {
		return b
	}
	if a > b {
		return b
	}
	return a
}

func read() []string {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = str[:len(str)-1]
	strs := strings.Split(str, " ")
	return strs
}
