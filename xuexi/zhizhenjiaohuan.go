package main

import "fmt"

func main() {
	a, b := 10, 20
	swap(&a, &b)
	fmt.Println("helloworld")
	fmt.Println(a, b)
}
func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
