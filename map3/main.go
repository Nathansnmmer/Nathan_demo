package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	//按照某个顺序来遍历map
	var coreMap = make(map[string]int, 100)
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100) //rand math 可以生成0~99的随机数
		coreMap[key] = value
	}
	//for k, v := range coreMap {
	//	fmt.Println(k, v)
	//}
	// 1 ,先创建一个切片将coremap中的key放入切片中
	// 2 ，然后对放入切片中的key进行排序
	// 3 ，然后按照排序好的切片中的key对map中的值排序
	keys := make([]string, 0, 100)
	for k := range coreMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	//切片中第一个是索引 “_” 不需要索引
	for _, key := range keys {
		fmt.Println(key, coreMap[key])
	}
}
