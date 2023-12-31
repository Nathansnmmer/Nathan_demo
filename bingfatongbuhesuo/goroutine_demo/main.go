package main

//互斥锁
import (
	"fmt"
	"sync"
)

// 用多个goroutine来并发操作全局变量x
var (
	x    int
	wg   sync.WaitGroup
	lock sync.Mutex //互斥锁
)

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock() //加锁
		x += 1
		lock.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
