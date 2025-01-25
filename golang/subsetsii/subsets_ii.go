/**
 * 90 - Medium
 *
 * Given an integer array nums that may contain duplicates, return all possible subsets (the power set).
 * The solution set must not contain duplicate subsets. Return the solution in any order.
 *
 * Example 1:
 *
 * 	Input: nums = [1,2,2]
 * 	Output: [[],[1],[1,2],[1,2,2],[2],[2,2]]
 *
 * Example 2:
 *
 * 	Input: nums = [0]
 * 	Output: [[],[0]]
 *
 * Constraints:
 *
 * 	- 1 <= nums.length <= 10
 * 	- -10 <= nums[i] <= 10
 */
package subsetsii

import "slices"

func subsetsWithDup(nums []int) [][]int {
	subsets := [][]int{}
	slices.Sort(nums)

	backtrack(&subsets, nums, []int{})

	return subsets
}

func backtrack(subsets *[][]int, nums, subset []int) {
	new := make([]int, 0, len(subset))
	new = append(new, subset...)
	*subsets = append(*subsets, new)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		
		subset = append(subset, nums[i])
		backtrack(subsets, nums[i+1:], subset)
		subset = subset[:len(subset)-1]
	}
}
