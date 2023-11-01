package main

//goroutine
import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello() {
	fmt.Println("hello 鸡虚")
	wg.Done() //通知wg把计数牌减一
}
func main() { //main函数会开启一个主goroutine去执行main函数
	wg.Add(1)  //计数牌  代表启动几个goroutine
	go hello() //开启了一个独立的goroutine去执行hello函数
	fmt.Println("hello main")
	//time.Sleep(time.Second) 等待俩goroutine都做完工作 但是这个sleep不建议使用它占用cpu太生硬了
	// main函数是老大（主goroutine）它没了（运行完了）老大手下的小弟（就是go hello 中手动添加的goroutine）就没了
	wg.Wait() //阻塞，等所有小弟都做完收兵  代表goroutine都做完了
}
