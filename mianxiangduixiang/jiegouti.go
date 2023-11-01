package main

import "fmt"

type Person struct {
	name  string
	age   int
	sex   string
	fight int
}
type Sperman struct {
	strength int
	speed    int
	p        Person
}

func (p *Person) setAge(age int) {
	p.age = age
}
func (s *Sperman) print() {
	fmt.Printf("%+v\n", s)
}
func main() {
	p1 := Person{"战五渣", 30, "man", 5}
	s1 := Sperman{
		strength: 10000,
		speed:    12,
		p:        p1,
	}
	s1.p.setAge(40)
	s1.print()
}
