package main

import (
	"fmt"
)

func main() {
	// 定义一个元素类型为map的切片
	//var smap = make([]map[string]int, 8, 8) //切片初始化了但是map还没有初始化
	////初始化map才能添加键值对
	//smap[0] = make(map[string]int, 8)
	//smap[0]["gouziji"] = 100
	//fmt.Println(smap)
	//定义一个值是切片的map
	//var nmap = make(map[string][]int, 8)前后这俩代码这个8可以省略
	var nmap = make(map[string][]int, 8)
	v, ok := nmap["jixu"]
	if ok {
		fmt.Println(v)
	} else {
		nmap["jixu"] = make([]int, 8)
		nmap["jixu"][0] = 100
		nmap["jixu"][1] = 100
		nmap["jixu"][2] = 100
	}
	for k, v := range nmap {
		fmt.Println(k, v)
	}
}
