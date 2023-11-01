package main

import "fmt"

func findTwoSum(arr []int, target int) [][]int {
	result := [][]int{}

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				result = append(result, []int{i, j})
			}
		}
	}

	return result
}

func main() {
	arr := []int{1, 3, 5, 7, 8}
	target := 8

	indices := findTwoSum(arr, target)

	if len(indices) > 0 {
		fmt.Printf("和为 %d 的两个元素的下标是：\n", target)
		for _, indexPair := range indices {
			fmt.Printf("(%d, %d)\n", indexPair[0], indexPair[1])
		}
	} else {
		fmt.Println("没有找到和为", target, "的两个元素。")
	}
}
