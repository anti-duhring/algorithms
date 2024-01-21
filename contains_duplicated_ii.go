package main

import (
	"fmt"
)

func main() {

	result := containsNearbyDuplicate([]int{1, 2, 3, 1}, 3)
	fmt.Printf("Result: %v\n", result)
}

func containsNearbyDuplicate(nums []int, k int) bool {
	window := []int{}

	for i := 0; i < len(nums); i++ {
		if i > k {
			// If the window size becomes larger than k we remove the first element
			window = window[1:]
		}
		if has(window, nums[i]) {
			// If the window already has this same value, so the result is true
			return true
		}
		window = append(window, nums[i])
	}
	return false
}

func has(arr []int, el int) bool {
	for _, v := range arr {
		if v == el {
			return true
		}
	}

	return false
}
