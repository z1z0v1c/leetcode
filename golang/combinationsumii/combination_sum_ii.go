/**
 * 40 - Medium
 *
 * Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in
 * candidates where the candidate numbers sum to target.
 *
 * Each number in candidates may only be used once in the combination.
 *
 * Note: The solution set must not contain duplicate combinations.
 *
 *
 * Example 1:
 *
 * 	Input: candidates = [10,1,2,7,6,1,5], target = 8
 * 	Output: [[1,1,6], [1,2,5], [1,7], [2,6]]
 *
 * Example 2:
 *
 * 	Input: candidates = [2,5,2,1,2], target = 5
 * 	Output: [[1,2,2], [5]]
 *
 * Constraints:
 *
 * 	- 1 <= candidates.length <= 100
 * 	- 1 <= candidates[i] <= 50
 * 	- 1 <= target <= 30
 */
package combinationsumii

import "slices"

func combinationSum2(candidates []int, target int) [][]int {
	sums := [][]int{}
	slices.Sort(candidates)

	backtrack(&sums, []int{}, candidates, target, 0)

	return sums
}

func backtrack(sums *[][]int, temp, candidates []int, target, start int) {
	if target < 0 {
		return
	} else if target == 0 {
		// Alocate new mwmory to avoid bugs
		new := make([]int, 0, len(temp))
		new = append(new, temp...)
		
		*sums = append(*sums, new)
	} else {
		for i := start; i < len(candidates); i++ {
			if candidates[i] > target {
				break
			}

			if i > start && candidates[i] == candidates[i-1] {
				continue
			}

			temp = append(temp, candidates[i])

			backtrack(sums, temp, candidates, target-candidates[i], i+1)

			temp = temp[:len(temp)-1]
		}
	}
}