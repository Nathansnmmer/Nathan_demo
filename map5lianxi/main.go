package main

import (
	"fmt"
	"strings"
)

// 练习判断map中how do you do 中的三个单词出现了几次
func main() {
	var s = "how do you do"
	var wordmap = make(map[string]int, 10)
	words := strings.Split(s, " ")
	for _, word := range words {
		v, ok := wordmap[word]
		if ok {
			wordmap[word] = v + 1
		} else {
			wordmap[word] = 1
		}
	}
	for k, v := range wordmap {
		fmt.Println(k, v)
	}
}
