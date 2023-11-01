package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("begin call goroutine")
	go func() {
		fmt.Println("i am a goroutine")
	}()
	time.Sleep(time.Second * 1)
	fmt.Println("end call goroutine")
}
