package main

import "fmt"

func main() {
	a2 := [3][4]int{
		{1, 2, 3, 4},
		{10, 40, 30, 20},
		{11, 12, 23, 34},
	}
	fmt.Println(a2)
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("i=%d, j=%d, val=%d\n", i, j, a2[i][j])
		}
	}
	a1 := []int{10, 20, 30, 40}
	for _, v := range a1 {
		fmt.Println(v)
	}
}
