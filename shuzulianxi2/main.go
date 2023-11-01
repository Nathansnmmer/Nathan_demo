package main

import (
	"fmt"
	"sort"
)

func main() {
	var a = [...]int{3, 7, 8, 9, 1}
	sort.Ints(a[:])
	fmt.Println(a)
}
