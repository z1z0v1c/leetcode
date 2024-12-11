/*
5 - Medium

Given a string s, return the longest palindromic substring in s.

Example 1:

	Input: s = "babad"
	Output: "bab"
	Explanation: "aba" is also a valid answer.

Example 2:

	Input: s = "cbbd"
	Output: "bb"

Constraints:

	- 1 <= s.length <= 1000
	- s consist of only digits and English letters.
*/
package longestpalindromicsubstring

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	lp := string(s[0])

	for i := 0; i < len(s); i++ {
		for j := len(s); j > i; j-- {
			sub := string(s[i:j])
			// Skip checking if the substring is shorter
			// than the current longest palindrome
			if len(sub) <= len(lp) {
				break
			}
			if isPalindrome(&sub) {
				if len(sub) > len(lp) {
					lp = sub
				}
			}
		}
	}

	return lp
}

// Pass by reference to improve performance, some tests are exhausting
func isPalindrome(s *string) bool {
	for i := 0; i < len(*s)/2; i++ {
		if (*s)[i] != (*s)[len(*s)-i-1] {
			return false
		}
	}
	return true
}

