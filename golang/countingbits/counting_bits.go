/*
*
  - 338 - Easy
    *
  - Given an integer n, return an array ans of length n + 1 such that for each i (0 <= i <= n), ans[i] is the number of 1's in the binary representation of i.

Example 1:

	Input: n = 2
	Output: [0,1,1]
	Explanation:

		0 --> 0
		1 --> 1
		2 --> 10

Example 2:

	Input: n = 5
	Output: [0,1,1,2,1,2]
	Explanation:

		0 --> 0
		1 --> 1
		2 --> 10
		3 --> 11
		4 --> 100
		5 --> 101

Constraints:

  - 0 <= n <= 10^5

Follow up:

	It is very easy to come up with a solution with a runtime of O(n log n).
	Can you do it in linear time O(n) and possibly in a single pass?
	Can you do it without using any built-in function (i.e., like __builtin_popcount in C++)?
*/
package countingbits

func countBits(n int) []int {
	counts := []int{}

	for i := 0; i <= n; i++ {
		counts = append(counts, 0)
		num := i

		for num > 0 {
			num &= num - 1
			counts[i]++
		}
	} 

	return counts
}

/**
 * ======> Bitwise operations solution <======
 *
 * func countBits2(n int) []int {
 * 	counts := []int{}
 * 
 * 	for i := 0; i <= n; i++ {
 * 		counts = append(counts, 0)
 * 		num := i
 * 
 * 		for num > 0 {
 * 			counts[i] += num & 1
 * 			num >>= 1
 * 		}
 * 	} 
 * 
 * 	return counts
 * }
 */

/**
 * ======> math/bits OnesCount solution <======
 *
 * import "math/bits"
 * 
 * func countBits(n int) []int {
 * 	counts := []int{}
 * 
 * 	for i := 0; i <= n; i++ {
 * 		counts = append(counts, bits.OnesCount(uint(i)))
 * 	} 
 * 
 * 	return counts
 * }
 */
