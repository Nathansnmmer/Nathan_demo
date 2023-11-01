package main

import "fmt"

/*
两个goroutine 两个channel

	1.生成0-100的数字发送给ch1

	2.从ch1中取出数据计算它的平方把结果发送到ch2中
*/
//生成0-100的数字发送给ch1
func f1(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

// 从ch1中取出数据计算它的平方把结果发送到ch2中
func f2(c1 <-chan int, c2 chan<- int) {
	//从通道中取值的方式一
	for {
		tmp, ok := <-c1
		if !ok {
			break
		}
		c2 <- tmp * tmp
	}
	close(c2)
}
func main() {
	ch1 := make(chan int, 2)
	ch2 := make(chan int, 2)
	go f1(ch1)
	go f2(ch1, ch2)
	//从通道中取值的方式二
	for ret := range ch2 {
		fmt.Println(ret)
	}
}
