// Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes
// the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.
// Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

// Example 1:

// Input: x = 123
// Output: 321

// Example 2:

// Input: x = -123
// Output: -321

// Example 3:

// Input: x = 120
// Output: 21

// Constraints:

// -231 <= x <= 231 - 1

package reverseinteger

import (
	"strconv"
)

func reverse(x int) int {
	xs := strconv.Itoa(x)

	var rsd string
	if x < 0 {
		rsd += string(xs[0])
		for i := 1; i < len(xs); i++ {
			rsd += string(xs[len(xs)-i])
		}
	} else {
		for i := range xs {
			rsd += string(xs[len(xs)-1-i])
		}
	}

	reversed, err := strconv.Atoi(rsd)
	if err != nil {
		println(err)
	}

	return reversed
}
