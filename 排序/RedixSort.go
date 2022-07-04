package main

import "fmt"

// 基排序
// 按照位数遍历
// 每次相同位的放在同一队列中排序

// RadixSort nums 数组，number 是该数组中数据的最大长度
func RadixSort(nums []int, number int) {
	// 开辟一个二维数组挂载数据
	var arr [10][]int
	var base = 1
	for k := 0; k < number; k++ {
		// 将nums中的数据按照 k 位放入 arr 中
		for j := 0; j < len(nums); j++ {
			flag := nums[j] / base % 10
			arr[flag] = append(arr[flag], nums[j])
		}
		base *= 10
		// 将 arr 中的数据依次取出放到 nums 中
		i := 0
		for n := 0; n < 10; n++ {
			for len(arr[n]) > 0 {
				nums[i] = arr[n][0]
				arr[n] = arr[n][1:]
				i++
			}
		}
	}
}

// 将 n 个身份证号排序
func main() {
	arr := []int{
		412314194511051354,
		422314197511051354,
		432314194511051354,
		442314194511051354,
		452314194511051354,
		462314194511051354,
		472314194511051354,
		492314194511051354,
		482314194511051354,
		380783192104031753,
		380783192104031754,
		380783192104021755,
	}
	RadixSort(arr, 18)
	for _, v := range arr {
		fmt.Println(v)
	}
}
