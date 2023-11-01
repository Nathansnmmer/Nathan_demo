package main

import "fmt"

func main() {
	var s1 []int
	s1 = append(s1, 3, 4, 5, 1)
	printSilce(s1)
	s2 := make([]int, 3)
	printSilce(s2)
	s2 = append(s2, 5)
	printSilce(s2)
	s1 = append(s1, 10)
	printSilce(s1)
}
func printSilce(s []int) {
	fmt.Printf("len=%d,cap=%d,s=%v\n", len(s), cap(s), s)
}
