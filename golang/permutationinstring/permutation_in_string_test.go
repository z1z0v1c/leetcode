package permutationinstring

import "testing"

func TestCheckInclusion(t *testing.T) {
	// Example one
	s1, s2 := "ab", "eidbaooo"

	if !checkInclusion(s1, s2) {
		t.Errorf("checkInclusion(%s, %s) should return true.", s1, s2)
	}

	// Example two
	s1, s2 = "ab", "eidboaoo"

	if checkInclusion(s1, s2) {
		t.Errorf("checkInclusion(%s, %s) should return false.", s1, s2)
	}

	// Example three
	s1, s2 = "adc", "dcda"

	if !checkInclusion(s1, s2) {
		t.Errorf("checkInclusion(%s, %s) should return true.", s1, s2)
	}

	// Example four
	s1, s2 = "ab", "a"

	if checkInclusion(s1, s2) {
		t.Errorf("checkInclusion(%s, %s) should return false.", s1, s2)
	}
}