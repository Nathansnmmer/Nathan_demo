package main

import "fmt"

// 闭包=函数+外层变量的引用
// 闭包的返回值是一个函数
func a(name string) func() {
	return func() {
		fmt.Println("hello", name)
	}
}

func main() {
	r := a("狗鸡虚") //此时r就是一个闭包
	r()
}
