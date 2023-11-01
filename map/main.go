package main

import "fmt"

// map
// map 不声明初始值就是nil就是空的
// map是键 值 对
func main() {
	var a map[string]int
	//map的初始化
	a = make(map[string]int, 8)
	fmt.Println(a == nil)
	//map中添加键 值 对
	a["guojix"] = 100
	a["gouzixu"] = 200
	fmt.Printf("a:%#v\n", a)
	fmt.Printf("type:%T\n", a)
	//map和变量一样可以同时声明和初始化
	b := map[int]bool{
		1: true,
		2: false,
	}
	fmt.Printf("b:%#v\n", b)
	fmt.Printf("type:%T\n", b)
}
