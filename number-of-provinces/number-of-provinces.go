package main

import "fmt"

func main() {

	given := [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	expect := 3
	result := findCircleNum(given)

	if result != expect {
		fmt.Printf("expected %v and got %v \n", expect, result)
		return
	}
	fmt.Println("test passed!!")
}

func findCircleNum(isConnected [][]int) int {
	provinces := len(isConnected)

	for i, v := range isConnected {

		for i2, v2 := range v {
			if v2 == 1 && i2 != i {
				provinces++
			}
		}
	}

	return provinces
}
