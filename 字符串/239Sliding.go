package main

import "fmt"

// 单调队列和优先级队列区别
// 单调队列，保证每次取的是队列中最大值
// 优先级队列，保证队列中的数据是有序的

// 思路设计一个单调队列
// 需求每次取的数据都是最大的，可以放入数据

type Queue struct {
	arr []int
}

// 获取最大值
func (q *Queue) getMax() int {
	return q.arr[0]
}

func (q *Queue) isempty() bool {
	return len(q.arr) == 0
}

// 入队列
// 大于哪个数哪个数后的全部舍弃
func (q *Queue) push(number int) {
	for !q.isempty() && number >= q.getMax() {
		q.arr = q.arr[:len(q.arr)-1]
	}
	q.arr = append(q.arr, number)
}

// 出队列,如果最大的出了，则更新队列，否则不更新
func (q *Queue) pop(number int) {
	if !q.isempty() && number >= q.arr[0] {
		q.arr = q.arr[1:]
	}
}

func main() {
	nums := []int{1, -1}
	k := 1
	fmt.Println(maxSlidingWindow(nums, k))
}

func maxSlidingWindow(nums []int, k int) []int {
	var start, end = 0, 0
	var queue = Queue{}
	var result []int
	for end = range nums {
		queue.push(nums[end])
		if end-start+1 == k {
			result = append(result, queue.getMax())
			queue.pop(nums[start])
			start++
		}
	}
	return result
}
