package 排序

/*
整体思路：
1、首先调整为小跟堆：从根节点依次遍历，将左右子节点最大值与父节点交换
2、循环遍历将尾节点与根节点交换。交换首尾节点，交换后将该节点依次往下调整为小跟堆。

小根堆的根节点小于左右子节点的值。
向下调整，假设已经是小跟堆了，插入或者删除一个数据时，需要将该数据节点及一下调整一次，调整为小跟堆
由于首尾交换小跟堆能够排序为降序。大根堆排序为升序。
*/

// HeapSortDES 降序
func HeapSortDES(arr []int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		adjustNodeDES(arr, i, len(arr))
	}

	for j := len(arr) - 1; j > 0; j-- {
		swap(arr, 0, j)
		adjustNodeDES(arr, 0, j)
	}
}

// HeapSortASC 升序
func HeapSortASC(arr []int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		adjustNodeASC(arr, i, len(arr))
	}

	for j := len(arr) - 1; j > 0; j-- {
		swap(arr, 0, j)
		adjustNodeASC(arr, 0, j)
	}
}

func swap(arr []int, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}

func adjustNodeASC(arr []int, nodeNum int, length int) {
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

func adjustNodeDES(arr []int, nodeNum int, length int) {
	temp := arr[nodeNum]
	for i := nodeNum*2 + 1; i < length; i = i*2 + 1 {
		if i+1 < length && arr[i] > arr[i+1] {
			i++
		}
		if arr[i] < temp {
			arr[nodeNum] = arr[i]
			nodeNum = i
		} else {
			break
		}
	}
	arr[nodeNum] = temp
}
