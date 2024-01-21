package main

import (
	"fmt"
)

func main() {
	result := lengthOfLongestSubstring("abcabcbb")

	fmt.Println(result)

	result = lengthOfLongestSubstring("bbbbb")

	fmt.Println(result)

	result = lengthOfLongestSubstring("aab")

	fmt.Println(result)
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 1 {
		return 1
	}

	maxLength := 0

	start := 0
	end := 0
	currentMap := []string{}

	for end < len(s) {
		repeatedString := false
		for i, val := range currentMap {
			if val == string(s[end]) {
				repeatedString = true
				start = i
			}
		}
		currentMap = append(currentMap, string(s[end]))

		if repeatedString == true {
			if maxLength < len(currentMap) {
				maxLength = len(currentMap)
			}
			currentMap = currentMap[start+1:]
		}

		fmt.Printf("string: %v, isDiff: %v\n", currentMap, repeatedString)
		end = end + 1
	}

	if maxLength == 0 {
		return len(currentMap)
	}

	return maxLength
}
