/**
 * 647 - Medium
 *
 * Given a string s, return the number of palindromic substrings in it.
 * A string is a palindrome when it reads the same backward as forward.
 * A substring is a contiguous sequence of characters within the string.
 *
 * Example 1:
 *
 * 	Input: s = "abc"
 * 	Output: 3
 * 	Explanation: Three palindromic strings: "a", "b", "c".
 *
 * Example 2:
 *
 * 	Input: s = "aaa"
 * 	Output: 6
 * 	Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".
 *
 * Constraints:
 *
 * 	- 1 <= s.length <= 1000
 * 	- s consists of lowercase English letters.
 */
package palindromicsubstrings

func countSubstrings(s string) int {
	count := 0

	mem := make([][]bool, len(s))

	for i := 0; i < len(s); i++ {
		row := make([]bool, len(s))

		row[i] = true
		mem[i] = row

		count++
	}

	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ {
			if isPalindrome(&mem, s, i, j) {
				count++
			}
		}
	}

	return count
}

func isPalindrome(mem *[][]bool, s string, i, j int) bool {
	if j-i <= 1 && s[i] == s[j] {
		(*mem)[i][j] = true
		return true
	}

	if (*mem)[i+1][j-1] && s[i] == s[j] {
		(*mem)[i][j] = true
		return true
	}

	return false
}
