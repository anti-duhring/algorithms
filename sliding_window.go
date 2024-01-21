package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	k := 3
	result := slidingWindow(arr, k)
	fmt.Printf("result %v\n", result)

	fmt.Println(arr)
	result2 := slidingWindowDynamic(arr, 7)
	fmt.Printf("result %v\n", result2)
}

func slidingWindow(arr []int, k int) []int {
	if k > len(arr) {
		return nil
	}
	slide := []int{}

	res := 0
	for i := 0; i < k; i++ {
		res += arr[i]
	}
	slide = append(slide, res)

	for i := 1; i < len(arr)-k+1; i++ {
		res = res - arr[i-1]
		res = res + arr[i+k-1]
		slide = append(slide, res)
	}

	return slide
}

func slidingWindowDynamic(arr []int, x int) int {
	minLength := len(arr)

	start := 0
	end := 0
	currentSum := 0

	for end < len(arr) {
		currentSum = currentSum + arr[end]
		end = end + 1

		for start < end && currentSum >= x {
			currentSum = currentSum - arr[start]
			start = start + 1
			if end-start+1 < minLength {
				minLength = end - start + 1
			}
		}
	}

	return minLength
}
