package 排序

// FindKthLargest 先调整为大根堆
// 然后交换首尾k次，第k次，这里注意调整完大根堆后，根节点为最大值，交换后最大值，到数组尾部。首尾交换k次后，第k大的数在数组length-k的位置
func FindKthLargest(nums []int, k int) int {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		adjustNodeASC1(nums, i, len(nums))
	}
	for j := len(nums) - 1; j > len(nums)-1-k; j-- {
		swap(nums, 0, j)
		adjustNodeASC1(nums, 0, j)
	}
	return nums[len(nums)-k]
}

func adjustNodeASC1(arr []int, nodeNum int, length int) {
	temp := arr[nodeNum]
	for i := nodeNum*2 + 1; i < length; i = i*2 + 1 {
		if i+1 < length && arr[i] < arr[i+1] {
			i++
		}
		if arr[i] > temp {
			arr[nodeNum] = arr[i]
			nodeNum = i
		} else {
			break
		}
	}
	arr[nodeNum] = temp
}
