package main

import (
	"../排序"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
输入：1 2 4 5 6 7
输出：
*/

func main() {
	reader := bufio.NewReader(os.Stdin)
	num, _ := reader.ReadString('\n')
	num = num[:len(num)-1]
	nums := strings.Split(num, " ")
	var arr []int
	for i := 0; i < len(nums); i++ {
		number, _ := strconv.Atoi(nums[i])
		arr = append(arr, number)
	}
	fmt.Println("排序前为", arr)
	排序.HeapSortDES(arr)
	fmt.Println("降序后为", arr)
	排序.HeapSortASC(arr)
	fmt.Println("升序后：", arr)
}
