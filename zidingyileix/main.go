package main

import "fmt"

// 自定义类型和类型别名
// 1 type 定义自定义类型
// MyInt 大写是暴露给外部的 它基于int类型的自定义类型
type myint int

// 类型别名
//newint int类型别名

type newint = int

func main() {
	var i myint
	fmt.Printf("type:%T value:%v\n", i, i)
	var r newint
	fmt.Printf("type:%T value:%v\n", r, r)
}
