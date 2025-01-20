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

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	var fq1, fq2 [26]int

	for i := 0; i < len(s1); i++ {
		fq1[s1[i]-'a']++
		fq2[s2[i]-'a']++
	}

	if fq1 == fq2 {
		return true
	}


	for i := len(s1); i < len(s2); i++ {
		fq2[s2[i]-'a']++
		fq2[s2[i-len(s1)]-'a']--

		if fq1 == fq2 {
			return true
		}
	}

	return false
}
