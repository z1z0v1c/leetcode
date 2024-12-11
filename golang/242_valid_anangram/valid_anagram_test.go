package validanagram

import "testing"

func TestIsAnagram(t *testing.T) {
	// Example one
	s1, s2 := "anagram", "nagaram"
	result := isAnagram(s1, s2)

	if !result {
		t.Errorf("testAnagram(%s, %s) should return true.", s1, s2)
	}

	// Example two
	s1, s2 = "rat", "car"
	result = isAnagram(s1, s2)

	if result {
		t.Errorf("testAnagram(%s, %s) should return false.", s1, s2)
	}
}