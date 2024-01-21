package main

import "fmt"

func main() {
	result := linearSearch([]int{1, 2, 3, 4}, 7)
	fmt.Println(result)
}

func linearSearch(arr []int, val int) bool {
	for v := range arr {
		if v == val {
			return true
		}
	}

	return false
}

