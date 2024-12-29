/* 
680 - Easy

Given a string s, return true if the s can be palindrome after deleting at most one character from it.

Example 1:

	Input: s = "aba"
	Output: true

Example 2:

	Input: s = "abca"
	Output: true
	Explanation: You could delete the character 'c'.

Example 3:

	Input: s = "abc"
	Output: false

Constraints:

	- 1 <= s.length <= 105
	- s consists of lowercase English letters.
*/
package validpalindromeii

func validPalindrome(s string) bool {
	start := 0
	end := len(s) - 1

	for start < end {
		if s[start] != s[end] {
			return isPalindrome(s[start+1:end+1]) || isPalindrome(s[start:end])
		}
		start++
		end--
	}

	return true
}

func isPalindrome(s string) bool {
	start := 0
	end := len(s) - 1

	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}

	return true
}

