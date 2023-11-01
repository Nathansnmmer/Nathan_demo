package main

import "fmt"

func main() {
	c := add_sub(10, 20, add)
	d := add_sub(30, 40, add)
	e := add_sub(30, 40, sub)
	fmt.Println(c, d, e)
}
func add_sub(a int, b int, f func(a int, b int) int) int {
	return f(a, b)
}
func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
