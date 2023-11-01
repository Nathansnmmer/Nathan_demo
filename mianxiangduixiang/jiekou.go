package main

import "fmt"

type Cat struct {
	color string
}

func (c *Cat) Sleep() {
	fmt.Println(c.color, "cat is sleeping")
}
func (c *Cat) Eating() {
	fmt.Println(c.color, "cat is eating")
}

type Dog struct {
	color string
}

func (d *Dog) Eating() {
	fmt.Println(d.color, "dog is eating")
}
func (d *Dog) Sleep() {
	fmt.Println(d.color, "dog is sleep")
}

type Animal interface {
	Sleep()
	Eating()
}

/*
	func main() {
		c1 := Cat{"write"}
		d1 := Dog{"black"}
		c1.Sleep()
		d1.Eating()
	}
*/
func main() {
	c1 := Cat{"write"}
	d1 := Dog{"balck"}
	var v1 Animal
	v1 = &c1
	v1.Eating()
	v1.Sleep()
	v1 = &d1
	d1.Eating()
	d1.Sleep()
}
