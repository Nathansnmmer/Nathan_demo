package main

import (
	"fmt"
)

func main() {
	c := func(a, b int) int {
		sum := a + b
		return sum
	}(10, 20)
	fmt.Println(c)
}
