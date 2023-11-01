package main

import "fmt"

// 方和和函数
// 方法是作用于一个特定类型变量的一个函数
// Person是一个结构体
type Person struct {
	name string
	age  int
}

// NewPerson是一个person的构造函数
func NewPerson(name string, age int) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

// Dream是为person定义的一个方法
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好go语言", p.name)
}

// setage是指针的修改年龄的方法
// 指针接收者指的就是接收者的类型是指针
func (p *Person) Setage(newage int) {
	p.age = newage
}

// setage2是一个使用值接受来修改年龄的方法
// 只有指针可以修改值是拷贝所以改不了
func (p Person) Setage2(newage int) {
	p.age = newage
}

func main() {
	p1 := NewPerson("武浩森", 18)
	// p1.Dream() //就等于*p1.Dream()
	fmt.Println(p1.age)
	p1.Setage(int(15))
	fmt.Println(p1.age)
	p1.Setage2(int(20))
	fmt.Println(p1.age)
}
