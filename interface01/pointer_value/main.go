package main

import "fmt"

//使用值接收者和使用指针接收者实现接口的区别

// 接口的嵌套
type animal interface { //animal接口可以实现mover接口和sayer接口里面的方法
	mover
	sayer
}

// 同一个类型可以实现不同的接口
type mover interface {
	move()
}
type sayer interface {
	say()
}
type person struct {
	name string
	age  int
}

// 值接收者实现接口：类型的值还是类型的指针都可以保存到接口变量中
//func (p person) move() {
//	fmt.Printf("%s在跑。。\n", p.name)
//}

// 使用指针接收者实现接口：只有类型的指针可以保存到接口变量中
func (p *person) move() {
	fmt.Printf("%s在跑。。\n", p.name)
}
func (p *person) say() {
	fmt.Printf("%s在叫 .. \n", p.name)
}
func main() {
	var m mover //定义一个mover类型的变量
	//p1 := person{ //p1是person类型的值
	//	name: "小鸡虚",
	//	age:  18,
	//}
	p2 := &person{ //p2是person类型的指针
		name: "狗鸡虚",
		age:  15,
	}
	//m = p1 //无法赋值p1是person类型的值没有实现mover接口    使用指针接收者实现接口  类型的值是不能赋值给接口变量的
	m = p2
	m.move()
	fmt.Println(m)
	var s sayer //定义一个sayer类型的变量
	s = p2
	s.say()
	fmt.Println(s)
	var a animal //定义一个animal类型的变量
	a = p2
	a.move()
	a.say()
	fmt.Println(a)
}
