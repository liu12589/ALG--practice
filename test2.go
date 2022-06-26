package main

import "fmt"

type coder1 interface {
	code()
	debug()
}

type Gopher struct {
	name string
}

// Go 语言是静态语言，严格的使用数据类型，即int和*int不是同一个种类型
// var c coder1 = Gopher{"lmy"} 实际上初始化的是一个 Gopher 值类型的接口变量
// 那么严格意义上来讲，只有 Gopher 值类型的接收者全部实现了接口中的函数，才能够初始乎
// 因此，*Gopher 指针类型实现的方法并不能归为 Gopher 值类型
// 那么严格意义上讲， Gopher 类型实现的方法并不能归为 *Gopher 类型。
// 那么 var c coder1 = &Gopher{"lmy"} 就会报错，而实际却不是
// 是因为在编译阶段，如果实现了接受者是值类型的方法，会隐含的也实现了接受者是指针类型的方法
func (g Gopher) code() {
	fmt.Println(g.name)
}

func (g *Gopher) debug() {
	fmt.Println(g.name)
}
func main() {
	var c coder1 = &Gopher{"lmy"}

	c.code()
	c.debug()
}
