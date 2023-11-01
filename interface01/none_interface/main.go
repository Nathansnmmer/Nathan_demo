package main

import (
	"fmt"
)

//空接口是没有定义任何方法的接口
//空接口可以存贮任意类型的变量
//空接口一般不需要提前定义

//空接口的应用
//1空接口类型作为函数的参数 比如println这个函数

//2空接口类型可以作为map的value 或者切片里面的类型

func main() {
	var x interface{}
	x = "hello"
	x = 18
	x = false
	fmt.Println(x)
	//var m = make(map[string]interface{}, 10)
	//m["name"] = "鸡虚"
	//m["age"] = 18
	//m["hobby"] = []string{"篮球", "足球"}
	//fmt.Println(m)
	ret, ok := x.(string) //类型断言 猜的类型不对时 ok=false ret =string类型的零值
	if !ok {
		fmt.Println("不是字符串类型")
	} else {
		fmt.Println("是字符串类型", ret)
	}
	//用swich进行类型断言
	switch v := x.(type) {
	case string:
		fmt.Printf("是字符串类型 value:%v\n", v)
	case bool:
		fmt.Printf("是布尔类型 value:%v\n", v)
	case int:
		fmt.Printf("是int类型 value:%v\n", v)
	default:
		fmt.Println("猜不到")
	}

}
