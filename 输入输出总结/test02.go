package main

import (
	"../排序"
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
		number, _ := strconv.Atoi(strs[i])
		arr = append(arr, number)
	}
	fmt.Println(arr)
	排序.QuickSort(arr)
	fmt.Println(arr)
}
