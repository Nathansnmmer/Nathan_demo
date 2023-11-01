package main

import "fmt"

// 结构体的匿名字段
type person struct {
	string
	int
}

func main() {
	p1 := person{
		"goujixu",
		18,
	}
	fmt.Println(p1)
	fmt.Println(p1.string)
	fmt.Println(p1.int)
}
