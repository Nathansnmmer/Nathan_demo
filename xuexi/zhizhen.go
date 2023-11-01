package main

import "fmt"

func main() {
	var a int = 10
	p := &a
	*p = 100
	fmt.Println(a, *p)
}
