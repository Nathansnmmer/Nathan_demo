package main

import (
	"fmt"
	"strings"
)

// 例子一
func makeSufix(sufix string) func(string) string {
	return func(name string) string {
		if !strings.HasPrefix(name, sufix) {
			return name + sufix
		}
		return name
	}
}

// 例子二
func cala(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}
func main() {
	r := makeSufix(".txt") //此时先给外层变量赋值 		此时r就是闭包
	rot := r("狗鸡虚")        //此时的r相当于内层的匿名函数再给匿名函数赋值
	fmt.Println(rot)
	x, y := cala(200)
	rot1 := x(100)
	fmt.Println(rot1)
	rot2 := y(100)
	fmt.Println(rot2)
}
