package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x      int
	wg     sync.WaitGroup
	rwlock sync.RWMutex // 读写互斥锁  就是读的操作远大于写适用于
)

func read() {
	rwlock.RLock()
	time.Sleep(time.Millisecond)
	rwlock.RUnlock()
	wg.Done()
}
func write() {
	rwlock.Lock()
	x += 1
	time.Sleep(time.Millisecond * 5)
	rwlock.Unlock()
	wg.Done()
}
func main() {
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
