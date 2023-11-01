package main

import "fmt"

// map的遍历用for rang 并且键值对的添加和顺序无关map中键值对添加时无序的
func main() {
	var a = make(map[string]int, 8)
	a["gouziji"] = 100
	a["guojixu"] = 200
	//map的删除是用一个内置的delet函数
	//删除gouziji这个键值对
	delete(a, "gouziji")
	for k, v := range a {
		fmt.Println(k, v)
	}
	//map也可以遍历一个值
	for k := range a {
		fmt.Println(k)
	}
	for _, v := range a {
		fmt.Println(v)
	}
}
