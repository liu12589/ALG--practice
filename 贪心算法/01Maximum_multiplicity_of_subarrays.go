package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b float64) float64 {
	if a > b {
		return b
	}
	return a
}

func maxMul(arr []float64) float64 {
	maxEnd, minEnd := arr[0], arr[0]
	maxNum, minNum := 0.0, 0.0
	res := 0.0
	for i := 0; i < len(arr); i++ {
		maxEnd = maxNum * arr[i]
		minEnd = minNum * arr[i]
		maxNum = max(max(arr[i], maxEnd), minEnd)
		minNum = min(min(arr[i], minEnd), maxEnd)
		res = max(res, maxNum)
	}
	return res
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// 读一行，去掉最后的换行后按空格分割
	str, _ := reader.ReadString('\n')
	str = str[:len(str)-1]
	strs := strings.Split(str, " ")

	// 将strs中字符串转位整数后传入
	n, _ := strconv.Atoi(strs[0])

	str, _ = reader.ReadString('\n')
	str = str[:len(str)-1]
	strs = strings.Split(str, " ")

	arr := make([]float64, n)

	for i := 0; i < n; i++ {
		arr[i], _ = strconv.ParseFloat(strs[i], 64)
	}

	fmt.Printf("%.2f", maxMul(arr))
}
