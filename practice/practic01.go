package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var arr []int

	str, _ := reader.ReadString('\n')
	str = str[:len(str)-1]
	strs := strings.Split(str, " ")

	for i := 0; i < len(strs); i++ {
		n, _ := strconv.Atoi(strs[i])
		arr = append(arr, n)
	}
	fmt.Println(arr)
	heapDES(arr)
	fmt.Println(arr)
}

// 利用堆排序降序
// 1、调整为小跟堆
// 2、遍历、交换首尾、调整为小跟堆
func heapDES(arr []int) {
	for i := len(arr)/2 - 1; i > 0; i-- {
		adjustHeap(arr, i, len(arr))
	}
	for i := len(arr) - 1; i > 0; i-- {
		swap(arr, 0, i)
		adjustHeap(arr, 0, i)
	}
}

func swap(arr []int, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}

func adjustHeap(arr []int, nodeNum int, length int) {
	temp := arr[nodeNum]
	for i := 2*nodeNum + 1; i < length; i = 2*i + 1 {
		if i+1 < length && arr[i] > arr[i+1] {
			i++
		}
		if temp > arr[i] {
			arr[nodeNum] = arr[i]
			nodeNum = i
		} else {
			break
		}
	}
	arr[nodeNum] = temp
}
