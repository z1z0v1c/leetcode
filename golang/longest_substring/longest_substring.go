package longest_substring

import (
	"strings"
)

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
