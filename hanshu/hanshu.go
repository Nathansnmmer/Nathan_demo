package main

import "fmt"

func main() {
	sum, sub := add_sub(1, 2)
	fmt.Println(sum, sub)

}
func add_sub(a, b int) (int, int) {
	sum0, sub0 := a+b, a-b
	return sum0, sub0
}
