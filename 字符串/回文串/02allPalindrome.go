package main

// 输出s所有的分割方案，
func main() {
	s := "google"

	judge := allIsPalindrome(s)

	//[[true false false true false false]
	// [false true true false false false]
	// [false false true false false false]
	// [false false false true false false]
	// [false false false false true false]
	//[false false false false false true]]
	for i := 0; i < len(judge); i++ {
		for j := i + 1; j < len(judge); j++ {
			if judge[i][j] {

			}
		}
	}

}

// 求所有的回文子串
func allIsPalindrome(str string) [][]bool {
	length := len(str)
	var isPalin = make([][]bool, length)
	for i := 0; i < length; i++ {
		isPalin[i] = make([]bool, length)
	}
	for i := 0; i < length; i++ {
		for h := 0; h < length; h++ {
			isPalin[i][h] = false
		}
	}
	var i, j, t int
	for t = 0; t < length; t++ {
		i = t
		j = t
		for i >= 0 && j < length && str[i] == str[j] {
			isPalin[i][j] = true
			i--
			j++
		}
		i = t
		j = t + 1
		for i >= 0 && j < length && str[i] == str[j] {
			isPalin[i][j] = true
			i--
			j++
		}
	}
	return isPalin
}
