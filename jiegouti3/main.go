package main

import "fmt"

// 嵌套结构体
type adders struct {
	city  string
	sheng string
}
type person struct {
	name   string
	age    int
	sex    string
	adders //嵌套另外一个结构体
}

func main() {
	p1 := person{
		name: "xijixu",
		age:  18,
		sex:  "女",
		adders: adders{
			city:  "hejin",
			sheng: "shanxi",
		},
	}
	fmt.Println(p1.name, p1.age, p1.sex, p1.adders)
	fmt.Println(p1.adders.sheng) //可以通过嵌套的匿名结构体访问其内部的字段
	fmt.Println(p1.sheng)        //也可以直接访问匿名结构体中的字段
}
