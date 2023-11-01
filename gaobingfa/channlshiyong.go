package main

import (
	"fmt"
	"sync"
	"time"
)

var w sync.WaitGroup
var c chan string

func reader() {
	msg := <-c
	fmt.Println("i am a reader", msg)
}
func main() {
	c = make(chan string)
	w.Add(1)
	go reader()
	fmt.Println("began sleep")
	c <- "hello"
	time.Sleep(time.Second * 1)
}
