/**
 * 131 - Medium
 *
 * Given a string s, partition s such that every substring of the partition is a palindrome.
 * Return all possible palindrome partitioning of s.
 *
 * Example 1:
 *
 * 	Input: s = "aab"
 * 	Output: [["a","a","b"],["aa","b"]]
 *
 * Example 2:
 *
 * 	Input: s = "a"
 * 	Output: [["a"]]
 *
 * Constraints:
 *
 * 	1 <= s.length <= 16
 * 	s contains only lowercase English letters.
 */
package palindromepartitioning

func partition(s string) [][]string {
	palindromes := make([][]string, 0)

	backtrack(s, &palindromes, &[]string{}, 0)

	return palindromes
}

func backtrack(s string, palindromes *[][]string, palindrome *[]string, start int) {
	if start == len(s) {
		copy := make([]string, 0, len(*palindrome))

		for i := 0; i < len(*palindrome); i++ {
			copy = append(copy, (*palindrome)[i])
		}

		*palindromes = append(*palindromes, copy)
		return
	}

	for i := start; i < len(s); i++ {
		if isPalindrome(s[start:i+1]) {
			*palindrome = append(*palindrome, s[start:i+1])
			backtrack(s, palindromes, palindrome, i+1)
			*palindrome = (*palindrome)[:len(*palindrome)-1]
		}
	} 

}

func isPalindrome(s string) bool {
	end := len(s) - 1

	for start := 0; start <= end; start++ {
		if s[start] != s[end] {
			return false
		}

		end--
	}

	return true
}
