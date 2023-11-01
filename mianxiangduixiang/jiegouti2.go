package main

import (
	"fmt"
)

type Ren struct {
	name  string
	age   int
	sex   string
	fight int
}
type Chaoman struct {
	strength int
	speed    int
	Ren
}

func (r *Ren) getAge(age int) {
	r.age = age
}
func (c *Chaoman) print() {
	fmt.Printf("%+v\n", c)
}
func main() {
	r1 := Ren{"zhanwuzha", 10, "man", 5}
	c1 := Chaoman{
		speed:    12,
		strength: 200,
		Ren:      r1,
	}
	c1.getAge(30)
	c1.print()
}
