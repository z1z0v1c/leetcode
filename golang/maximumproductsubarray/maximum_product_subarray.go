/**
 * 152 - Medium
 *
 * Given an integer array nums, find a subarray that has the largest product, and return the product.
 * The test cases are generated so that the answer will fit in a 32-bit integer.
 *
 * Example 1:
 *
 * 	Input: nums = [2,3,-2,4]
 * 	Output: 6
 * 	Explanation: [2,3] has the largest product 6.
 *
 * Example 2:
 *
 * 	Input: nums = [-2,0,-1]
 * 	Output: 0
 * 	Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
 *
 * Constraints:
 *
 * 	- 1 <= nums.length <= 2 * 10^4
 * 	- -10 <= nums[i] <= 10
 * 	- The product of any subarray of nums is guaranteed to fit in a 32-bit integer.
 */
package maximumproductsubarray

func maxProduct(nums []int) int {
	maxProduct := make([]int, len(nums))
	maxProduct[0] = nums[0]

	minProduct := make([]int, len(nums))
	minProduct[0] = nums[0]

	best := maxProduct[0]

	for i := 1; i < len(maxProduct); i++ {
		minProduct[i] = min(minProduct[i-1]*nums[i], nums[i])

		maxProduct[i] = max(maxProduct[i-1]*nums[i], minProduct[i-1]*nums[i])

		best = max(maxProduct[i], best)
	}

	return best
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
