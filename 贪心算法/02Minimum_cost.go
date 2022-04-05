package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type intHeap []int

func (ih intHeap) Len() int           { return len(ih) }
func (ih intHeap) Less(x, y int) bool { return ih[x] <= ih[y] }
func (ih intHeap) Swap(x, y int)      { ih[x], ih[y] = ih[y], ih[x] }
func (ih *intHeap) Push(x interface{}) {
	*ih = append(*ih, x.(int))
}
func (ih *intHeap) Pop() interface{} {
	old := *ih
	n := len(old)
	x := old[n-1]
	*ih = old[0 : n-1]
	return x
}

func main() {
	var m int
	fmt.Scanf("%d\n", &m)
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	cols := strings.Split(str[:len(str)-1], " ")

	arr := make(intHeap, 0)
	for _, v := range cols {
		num, _ := strconv.Atoi(v)
		arr = append(arr, num)
	}
	heap.Init(&arr)
	var result int

	for arr.Len() > 1 {
		num := heap.Pop(&arr).(int) + heap.Pop(&arr).(int)
		result += num
		heap.Push(&arr, num)
	}
	fmt.Println(result)
	return
}
