package 排序

/*
思路：选择排序列表中第一个记录为枢轴，在 r[0]位置上，附设两个指针 L、H
	1、H 指向最后一个元素，向左搜索。若 >= r[0],向左移动，< r[0]与 r[0] 交换。
	2、L 指向 r[1]，向右搜索。若 <= r[0],向右移动，否则与 r[0] 交换。
*/

func partition(arr []int, low int, high int) int {
	pivot := arr[low]
	for low < high {
		for low < high && arr[high] >= pivot {
			high--
		}
		arr[low] = arr[high]
		for low < high && arr[low] <= pivot {
			low++
		}
		arr[high] = arr[low]
	}
	arr[low] = pivot
	return low
}

func _quickSort(arr []int, low int, high int) {
	if low < high {
		pivot := partition(arr, low, high)
		_quickSort(arr, low, pivot-1)
		_quickSort(arr, pivot+1, high)
	}
}

func QuickSort(arr []int) {
	_quickSort(arr, 0, len(arr)-1)
}
