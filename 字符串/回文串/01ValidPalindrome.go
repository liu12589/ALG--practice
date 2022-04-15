package main

/*
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
说明：本题中，我们将空字符串定义为有效的回文串。
————————————————————————————————————————
剔除杂质、统一规格
回文判断
*/
import "fmt"

func main() {
	// 循环输出会参杂空字符
	arr := "A man, a plan, a canal: Panama"
	var noNullArr []byte
	for j := 0; j < len(arr); j++ {
		if (arr[j] <= 'z' && arr[j] >= 'a') || (arr[j] <= 'Z' && arr[j] >= 'A') || (arr[j] <= '9' && arr[j] >= '0') {
			if arr[j] <= 'Z' && arr[j] >= 'A' {
				noNullArr = append(noNullArr, arr[j]+32)
			} else {
				noNullArr = append(noNullArr, arr[j])
			}
		}
	}
	fmt.Println(palindrome(noNullArr))
}

func palindrome(arr []byte) bool {
	start := 0
	end := len(arr) - 1
	for start < end {
		if arr[start] != arr[end] {
			return false
		}
		start++
		end--
	}
	return true
}
