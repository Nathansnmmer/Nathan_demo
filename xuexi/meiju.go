package main

import "fmt"

const (
	lo = iota
	log
	user   = iota + 1
	accont = iota + 3
)
const (
	apple, banna = iota + 1, iota + 2
	peach, pear
	orange, mango
)

func main() {
	fmt.Println(apple, banna, peach, pear, orange, mango)
	fmt.Println(lo, log, user, accont)
}
