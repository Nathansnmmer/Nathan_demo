package main

import "fmt"

type person struct {
	name, city string
	age        int
}

// 结构体的构造函数   				但是实际上go语言中是没有构造函数的但是可以实现
func newPerson(name, city string, age int) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

func main() {
	r := newPerson("jixu", "bejing", 10)
	fmt.Printf("type:%T value:%#v\n", r, r)
}
