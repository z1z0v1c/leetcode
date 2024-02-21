package main

import (
	"fmt"
	"strings"
)

func main() {
	m := lengthOfLongestSubstring(" ")
	fmt.Println(m)

	m = lengthOfLongestSubstring("abcabcbb")
	fmt.Println(m)

	m = lengthOfLongestSubstring("bbbbb")
	fmt.Println(m)

	m = lengthOfLongestSubstring("pwwkew")
	fmt.Println(m)
}

func lengthOfLongestSubstring(s string) int {
	var max int

	for i, v := range s {
		tempMaxSub := string(v)

		for j := i + 1; j < len(s); j++ {
			if strings.Contains(tempMaxSub, string(s[j])) {
				break
			} else {
				tempMaxSub += string(s[j])
			}
		}

		if len(tempMaxSub) > max {
			max = len(tempMaxSub)
		}
	}

	return max
}
