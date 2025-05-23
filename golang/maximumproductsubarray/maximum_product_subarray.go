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
	maxProduct := nums[0]
	minProduct := nums[0]

	bestMaxProduct := nums[0]

	for i := 1; i < len(nums); i++ {
		tempMinProduct := minProduct

		minProduct = min(min(minProduct*nums[i], maxProduct*nums[i]), nums[i])
		maxProduct = max(max(maxProduct*nums[i], tempMinProduct*nums[i]), nums[i])

		bestMaxProduct = max(maxProduct, bestMaxProduct)
	}

	return bestMaxProduct
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
