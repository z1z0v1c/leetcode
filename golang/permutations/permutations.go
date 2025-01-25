/**
 * 46 - Medium
 *
 * Given an array nums of distinct integers, return all the possible permutations
 * You can return the answer in any order.
 *
 * Example 1:
 *
 * 	Input: nums = [1,2,3]
 * 	Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 *
 * Example 2:
 *
 * 	Input: nums = [0,1]
 * 	Output: [[0,1],[1,0]]
 *
 * Example 3:
 *
 * 	Input: nums = [1]
 * 	Output: [[1]]
 *
 * Constraints:
 *
 * 	- 1 <= nums.length <= 6
 * 	- -10 <= nums[i] <= 10
 * 	- All the integers of nums are unique.
 */
package permutations

func permute(nums []int) [][]int {
	permutations := [][]int{}

	backtrack(&permutations, nums, []int{})

	return permutations
}

func backtrack(permutations *[][]int, nums, permutation []int) {
	if len(nums) == 0 {
		// Allocate new memory to avoid bugs
		new := make([]int, 0, len(permutation))
		new = append(new, permutation...)
		
		*permutations = append(*permutations, new)
	} else {
		for i := 0; i < len(nums); i++ {
			permutation = append(permutation, nums[i])

			// Create new nums without used element
			newNums := make([]int, 0, len(nums)-1)
			newNums = append(newNums, nums[:i]...)
			newNums = append(newNums, nums[i+1:]...)

			backtrack(permutations, newNums, permutation)

			// Remove the last element after recursive call
			permutation = permutation[:len(permutation)-1]
		}
	}
}
