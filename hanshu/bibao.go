package main

import "fmt"

func main() {
	nextnumber := getSequence()
	fmt.Println(nextnumber())
	fmt.Println(nextnumber())
	fmt.Println(nextnumber())
	f := getSequence()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
