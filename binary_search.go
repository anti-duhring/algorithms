package main

import "fmt"

func main() {
	result := binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 6)
	fmt.Println(result)
}

func binarySearch(arr []int, target int) bool {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == target {
			return true
		}
		if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return false
}
