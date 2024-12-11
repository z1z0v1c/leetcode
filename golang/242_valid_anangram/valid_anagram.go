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

import (
	"sort"
	"strings"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	s1 := strings.Split(s, "")
	sort.Strings(s1)

	s2 := strings.Split(t, "")
	sort.Strings(s2)

	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}

	return true
}
