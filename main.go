package main

import "fmt"

func main() {
	val1 := []int{1, 3, 2, 3, 1, 3}
	val2 := 3
	result := longestEqualSubarray(val1, val2)
	fmt.Printf("result %v\n", result)
}

func longestEqualSubarray(nums []int, k int) int {
	dict := map[int][]int{}
	for i, num := range nums {
		dict[num] = append(dict[num], i)
	}

	fmt.Println(dict)

	ans := 1
	for _, arr := range dict {
		k1 := k
		// l := 0

		for r := 1; r < len(arr); r++ {
			k1 -= arr[r] - arr[r-1] - 1
			fmt.Printf("arr[r]: %v, arr[r-1]: %v, k: %v\n", arr[r], arr[r-1], arr[r]-arr[r-1]-1)
		}
	}

	return ans
}
