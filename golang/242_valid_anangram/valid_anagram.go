/* Easy

Given two strings s and t, return true if t is an anagram of s, and false otherwise.

Example 1:

	Input: s = "anagram", t = "nagaram"
	Output: true

Example 2:

	Input: s = "rat", t = "car"
	Output: false

Constraints:

	- 1 <= s.length, t.length <= 5 * 104
	- s and t consist of lowercase English letters.

Follow up: What if the inputs contain Unicode characters? How would you adapt your solution to such a case?
*/

package validanagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	runes1 := make(map[rune]int, len(s))
	for _, rn := range s {
		runes1[rn]++
	}

	runes2 := make(map[rune]int, len(t))
	for _, rn := range t {
		runes2[rn]++
	}

	for k, v := range runes1 {
		if v != runes2[k] {
			return false
		}
	}

	return true
}
