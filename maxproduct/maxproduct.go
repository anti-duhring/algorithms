package main

import (
	"fmt"
	"log"
)

func main() {

	max := maxProducts([]int{2, 2, 1, 8, 1, 5, 4, 5, 2, 10, 3, 6, 5, 2, 3})
	expect := 63

	if max != expect {
		log.Fatalf("expect %v and got %v", expect, max)
	} else {
		log.Fatal("test passed!!")
	}
}

func maxProducts(nums []int) int {
	j := nums[0]
	l := nums[1]

	for _, v := range nums[2:] {
		if v > j {
			if j > l {
				l = v
			} else {
				j = v
			}

			continue
		}

		if v > l {

			if l > j {
				j = v
			} else {
				l = v
			}
			continue
		}
	}

	fmt.Printf("max1: %v max2: %v \n", j, l)
	return (j - 1) * (l - 1)
}
