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
	str, _ := reader.ReadString('\n')
	str = str[:len(str)-1]
	strs := strings.Split(str, " ")
	length, _ := strconv.Atoi(strs[0])

	sum := 0
	for i := 0; i < len(strs)-1; i++ {
		number, _ := strconv.Atoi(strs[i+1])
		sum += number
	}
	fmt.Println((length+1)*length/2 - sum)
}
