package main

import "fmt"

type coder interface {
	showName()
	showAge()
	showSchool()
}
type Person struct {
	name string
	age  int
}

func (p Person) showName() {
	fmt.Println(p.age)
	fmt.Println(p.name)
}

func (p *Person) showAge() {
	fmt.Println(p.age)
	fmt.Println(p.name)
}

func (p *Person) showSchool() {
	fmt.Println("school")
}

func main() {
	var obj1 Person = Person{
		name: "lmy",
		age:  18,
	}
	var obj2 *Person = &Person{
		name: "cfy",
		age:  18,
	}
	var obj3 coder = &Person{
		name: "hzl",
		age:  10,
	}

	obj3.showName()
	obj3.showAge()
	obj3.showSchool()

	obj1.showName()
	obj2.showName()

	obj1.showAge()
	obj2.showAge()

}
