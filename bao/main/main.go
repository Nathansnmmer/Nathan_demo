package main

import (
	"fmt"
	heihei "main.go/add" //加入导入的包名字一样可以使用别名heihei这就是这个包的别名
)

func main() {
	heihei.Add(1, 2)
	fmt.Println(heihei.Name)
}
