/**
 * 3 - Medium
 * 
 * Given a string s, find the length of the longest substring without repeating characters.
 * 
 * Example 1:
 * 
 * 	Input: s = "abcabcbb"
 * 	Output: 3
 * 	Explanation: The answer is "abc", with the length of 3.
 * 
 * Example 2:
 * 
 * 	Input: s = "bbbbb"
 * 	Output: 1
 * 	Explanation: The answer is "b", with the length of 1.
 * 
 * Example 3:
 * 
 * 	Input: s = "pwwkew"
 * 	Output: 3
 * 	Explanation: The answer is "wke", with the length of 3.
 * 
 * 	Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
 * 
 * Constraints:
 * 
 * 	- 0 <= s.length <= 5 * 104
 * 	- s consists of English letters, digits, symbols and spaces.
 */
package longestsubstring

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

