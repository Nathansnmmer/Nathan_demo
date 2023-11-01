package main

import (
	"fmt"
	"time"
)

// work-pool
func worker(id int, job <-chan int, result chan<- int) {
	for j := range job {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		result <- j * 2
		time.Sleep(time.Millisecond * 500)
		fmt.Printf("worker:%d stop job:%d\n", id, j)
	}
}
func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 200)
	//开启三个goroutine
	for m := 0; m < 3; m++ {
		go worker(m, jobs, results)
	}
	//发送五个任务
	for i := 0; i < 5; i++ {
		jobs <- i
	}
	close(jobs)
	//输出结果
	for i := 0; i < 5; i++ {
		ret := <-results
		fmt.Println(ret)
	}
}
