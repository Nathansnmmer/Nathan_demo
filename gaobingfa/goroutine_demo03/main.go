package main

import (
	"fmt"
	"sync"
)

var ru sync.WaitGroup

func main() {
	ru.Add(100)
	for i := 0; i <= 100; i++ {
		go func(i int) {
			fmt.Println("hello", i)
			ru.Done()
		}(i)
	}
	fmt.Println("hello main")
	ru.Wait()
}
