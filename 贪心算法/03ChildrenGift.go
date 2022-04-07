package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var h, w []int
	h = readerArr()
	w = readerArr()
	maxChildren(h, w)
}

func maxChildren(h, w []int) {
	sum := 0
	sort.Ints(h)
	sort.Ints(w)
	for i := 0; i < len(w); i++ {
		if len(h) > 0 && w[i] >= h[0] {
			sum++
			h = h[1:]
		}
	}
	fmt.Println(sum)
}

func readerArr() []int {
	var h []int
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = str[:len(str)-1]
	strs := strings.Split(str, " ")
	hLength, _ := strconv.Atoi(strs[0])

	strH, _ := reader.ReadString('\n')
	strH = strH[:len(strH)-1]
	strHs := strings.Split(strH, " ")
	for i := 0; i < hLength; i++ {
		number, _ := strconv.Atoi(strHs[i])
		h = append(h, number)
	}
	return h
}
