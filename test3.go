package main

import "fmt"

type test interface {
	showArr()
}

type arrName struct {
	arr int
}

/*
接收者是值类型还是指针类型的区别，其实就是指向的是否是同一个区域，值类型传参时
直接复制，而指针类型会传指向该地址的指针
那么什么时候使用值类型，什么时候使用指针类型呢
值类型：静态变量、slice、map、interface、channel
指针类型：文件结构体，这种不能够被复制的又语言底层没有提供 header 的必须使用指针类型。
2
2
——————————————————————————————————————————————————
2
3
*/
func (n *arrName) showArr() {
	n.arr += 1
	fmt.Println(n.arr)
}
func main() {
	// 值类型是复制，不会修改对象
	var arr1 test = &arrName{1}
	test1(arr1)
	test2(arr1)
}

func test1(arr1 test) {
	arr1.showArr()
}

func test2(arr1 test) {
	arr1.showArr()
}
