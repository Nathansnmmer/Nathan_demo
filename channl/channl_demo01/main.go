package main

import "fmt"

func main() {
	//var ch1 chan int //引用类型 三种 切片  map 和channl 定义之后需要make初始化
	//ch1 = make(chan int, 1)
	ch1 := make(chan int, 1) //这段代码就等于上面俩句   有缓冲区 异步通道 	相当于把东西放到驿站
	//ch1 := make(chan int) //						无缓冲区又称为同步通道  必须手把手交给你
	ch1 <- 10
	x := <-ch1
	fmt.Println(x)
	//len(ch1) //拿到通道中元素的数量
	//cap(ch1) //拿到通道的容量		这俩用的比较少
	close(ch1)
}
