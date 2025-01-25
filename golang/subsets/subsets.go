/**
 * 78 - Medium
 *
 * Given an integer array nums of unique elements, return all possible subsets (the power set).
 * The solution set must not contain duplicate subsets. Return the solution in any order.
 *
 * Example 1:
 *
 * 	Input: nums = [1,2,3]
 * 	Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
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
 * 	- All the numbers of nums are unique.
 */
package subsets

import (
	"slices"
)

func subsets(nums []int) [][]int {
	sets := [][]int{{}}

	for _, num := range nums {
		length := len(sets)

		for j := 0; j < length; j++ {
			// Allocate new memory to avoid bugs
			new := make([]int, 0, len(sets[j])+1)

			new = append(new, sets[j]...)
			new = append(new, num)

			// Sort for testing
			slices.Sort(new)

			sets = append(sets, new)
		}
	}

	return sets
}
