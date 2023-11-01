package main

import (
	"fmt"
)

func main() {
	al := [5]int{1, 2, 3, 4, 5}
	sl := al[2:4]
	fmt.Println(al)
	fmt.Println(sl)
	sl[0] = 20
	fmt.Println("after-------")
	fmt.Println(al)
	fmt.Println(sl)
}
