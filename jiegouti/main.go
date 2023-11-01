package main

import "fmt"

//结构体

type person struct {
	name, city string
	age        int
}

func main() {
	/*	var p person
		p.city = "背景"
		p.name = "goujixu"
		p.age = 10
		fmt.Printf("%#v\n", p)
		fmt.Println(p.age)
		fmt.Println(p.city)
		fmt.Println(p.name)
		//匿名结构体 用在某些临时的复合的场景需要匿名结构体
		var jk struct {
			name  string
			where string
		}
		jk.where = "hejin"
		jk.name = "gouji"
		fmt.Printf("%#v\n", jk)

	*/
	//结构体的初始化 两种都是定义加初始化
	//1键值对
	//键值对赋值的时候可以某些字段不用写
	p1 := person{
		name: "jixu",
		age:  10,
	}
	p5 := &person{
		name: "xiaoguo",
		city: "beijing",
		age:  15,
	}
	fmt.Printf("%#v\n", p1)
	fmt.Printf("%#v\n", p5)
	//2 值的列表进行初始化 赋值的时候要按声明的时候自己按顺序来赋值
	p6 := person{
		"xiaowangzi",
		"shanxi",
		10,
	}
	fmt.Printf("%#v\n", p6)
}
