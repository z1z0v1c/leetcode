/**
 * 567 - Medium
 *
 * Given two strings s1 and s2, return true if s2 contains a permutation
 * of s1, or false otherwise.
 * 
 * In other words, return true if one of s1's permutations is the substring of s2.
 * 
 * Example 1:
 * 
 * 	Input: s1 = "ab", s2 = "eidbaooo"
 * 	Output: true
 * 	Explanation: s2 contains one permutation of s1 ("ba").
 * 
 * Example 2:
 * 
 * 	Input: s1 = "ab", s2 = "eidboaoo"
 * 	Output: false
 * 
 * Constraints:
 * 
 * 	- 1 <= s1.length, s2.length <= 10^4
 * 	- s1 and s2 consist of lowercase English letters.
 */
package permutationinstring

var fq1 map[rune]int

func checkInclusion(s1 string, s2 string) bool {
	fq1 = countLetters(s1)

	for start, end := 0, len(s1); end <= len(s2); end++ {
		if checkPermutation(s2[start:end]) {
			return true
		}

		start++
	}

	return false
}

func checkPermutation(s string) bool {
	fq := countLetters(s)

	for k, v := range fq {
		if fq1[k] != v {
			return false
		}
	}

	return true
}

func countLetters(s string) map[rune]int {
	fq := make(map[rune]int)

	for _, l := range s {
		fq[l]++
	}

	return fq
}
