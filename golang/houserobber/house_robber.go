/*
*
  - 198 - Medium
    *
  - You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed,

the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected and it
will automatically contact the police if two adjacent houses were broken into on the same night.

Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob
tonight without alerting the police.

Example 1:

	Input: nums = [1,2,3,1]
	Output: 4

	Explanation:

		- Rob house 1 (money = 1) and then rob house 3 (money = 3).
		- Total amount you can rob = 1 + 3 = 4.

Example 2:

	Input: nums = [2,7,9,3,1]
	Output: 12

	Explanation:

		- Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
		- Total amount you can rob = 2 + 9 + 1 = 12.

Constraints:

  - 1 <= nums.length <= 100
  - 0 <= nums[i] <= 400
*/
package houserobber

func Rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	mem := make([]int, len(nums)) 

	mem[0] = nums[0]
	mem[1] = nums[1]

	best, prevBest := Max(mem[0], mem[1]), 0

	for i := 2; i < len(nums); i++ {
		if best > mem[i-2] && best != mem[i-1] {
			mem[i] = nums[i] + best
		} else if prevBest > mem[i-2] {
			mem[i] = nums[i] + prevBest
		} else {
			mem[i] = nums[i] + mem[i-2]
		}

		if mem[i] >= best {
			prevBest = best
			best = mem[i]
		}
	}

	return best
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}