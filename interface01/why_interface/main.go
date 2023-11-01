package main

import "fmt"

// 为什么需要接口
type dog struct{}
type cat struct{}

// 不同的类型可以实现同一个接口
func (d dog) jiao() {
	fmt.Println("狗会汪汪汪")
}
func (c cat) jiao() {
	fmt.Println("猫会叫喵喵")
}

type person struct {
	name string
}

func (p person) jiao() {
	fmt.Println("啊啊啊")
}

// 接口不管你什么类型，只管实现方法
// 定义一个类型，一个抽象的类型，只要实现了jiao这个放的就叫sayer类型
type sayer interface {
	jiao()
}

func da(m sayer) {
	m.jiao() //不管传进来的是什么就打
}

func main() {
	//c1 := cat{}
	//da(c1)
	//d1 := dog{}
	//da(d1)
	//p1 := person{
	//	name: "小王子",
	//}
	//da(p1)
	var s sayer
	c1 := cat{}
	s = c1 //当你实现了这个接口 这个接口类型的变量就可以接受任意的只要你实现了接口的变量
	p1 := person{name: "小鸡虚"}
	s = p1
	fmt.Println(s)

}
