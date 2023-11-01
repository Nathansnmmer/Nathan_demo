package main

import "fmt"

func main() {
	p1 := Person{"战五渣", 15, 5, "男"}
	fmt.Println(p1)
	var p2 Person
	p2.name = "xiaohong"
	p2.age = 10
	p2.sex = "wuman"
	fmt.Printf("%+v\n", p2)
}

type Person struct {
	name  string
	age   int
	fight int
	sex   string
}
