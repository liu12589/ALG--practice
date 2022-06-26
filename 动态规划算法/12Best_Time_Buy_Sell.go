package main

import "fmt"

/*
题目1：已知后面N天每只股票的价格为P0、P1、P2、P3
可以最多买一支股票、卖一只股票
求最大利润
——————————————————————————————————————————————
x为当前最大利差，n为当前最小价格
———————————————————————————————————————————————
题目2：可以买卖任意多次，但任意时刻手中只能有一股，求最大利润
只要差大于0，就买卖
———————————————————————————————————————————————
题目3：可以最多进行两次买两次买，每次买卖都是一股
不能在卖光手中股票前买入，但可以同一天卖完后买入
分析（确定状态）：我们假设有四种状态（很容易想到）buy1、buy2、sell1、sell2。分别对应第1、2次买入卖出。如果我们知道第 i - 天结束后的四种状态，
我们怎么确定第 i 天的状态呢。
1、什么时候第一次买入？寻找第一个最小值 buy1 = min{buy1, prices[i]}
2、什么时候第一次卖出？寻找最大利差 gap1 = max{gap1, prices[i] - buy1}
3、什么时候第二次买入？寻找buy1后最大值 buy2 = min(sell1在数组中位置 <= i <= length）{buy2, prices[i]}
4、什么时候第二次卖出？寻找buy2后最大利差 sell2 = max(sell1在数组中位置 <= i <= length）{gap1, prices[i] - buy1}
*/
func main() {
	var arr = []int{7, 1, 5, 3, 6, 4}
	buyStock1(arr)
	//buyStock2(arr)
	buyStock3(arr)
}

func buyStock1(arr []int) {
	var x, n int // x为当前最大利差，n为当前最小价格

	x = 0
	n = arr[0]
	for i := 0; i < len(arr); i++ {
		gap := arr[i] - n
		if gap > x {
			x = gap
		}
		if arr[i] < n {
			n = arr[i]
		}
	}
	fmt.Println(x)
}

func buyStock2(arr []int) {
	res := 0
	for i := 0; i < len(arr); i++ {
		gap := arr[i+1] - arr[i]
		if gap > 0 {
			res += gap
		}
	}
	fmt.Println(res)
}

func buyStock3(arr []int) {
	var buy1 = arr[0]
	var gap1 = 0
	var buy2, gap2 = 0, 0
	for i := 1; i < len(arr); i++ {
		x1 := arr[i] - buy1
		x2 := arr[i] - buy2
		if gap1 < x1 {
			gap1 = x1
			gap2 = x2
		}
		if buy1 > arr[i] {
			buy1 = arr[i]
			buy2 = buy1
		}
	}
	fmt.Println(gap1 + gap2)
}
