package reversebits

import "testing"

func TestReverseBits(t *testing.T) {
	// Example one
	var n uint32 = 43261596
	var expected uint32 = 964176192
	actual := reverseBits(n)

	if expected != actual {
		t.Errorf("reverseBits(%d) returned %d insted of %d.", n, actual, expected)
	} 

	// Example two
	n = 4294967293
	expected = 3221225471
	actual = reverseBits(n)

	if expected != actual {
		t.Errorf("reverseBits(%d) returned %d insted of %d.", n, actual, expected)
	} 
}