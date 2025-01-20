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

	if !checkInclusion(s1, s2) {
		t.Errorf("checkInclusion(%s, %s) should return true.", s1, s2)
	}
}