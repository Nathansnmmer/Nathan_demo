package main

import "fmt"

// 结构体的继承
// 就是一个结构体中有一个嵌套的结构体这个结构体会继承嵌套结构体的方法也可以调用
type dog struct {
	name string
}

func (d *dog) jiao() {
	fmt.Printf("%s会叫\n", d.name)
}

type mao struct {
	jue int
	*dog
}

func (m mao) pao() {
	fmt.Printf("%s会跑\n", m.name)
}

func main() {
	d1 := mao{
		jue: 4,
		dog: &dog{
			name: "goujixu",
		},
	}
	d1.jiao()
	d1.pao()
}
