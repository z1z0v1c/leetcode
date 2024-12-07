package sortedarraytobst

import (
	cs "github.com/z1z0v1c/leetcode/commonstructs"
)

/*

108 -Easy

Given an integer array nums where the elements are sorted in ascending order, convert it to a
height-balanced binary search tree.

Example 1:

	Input: nums = [-10,-3,0,5,9]
	Output: [0,-3,9,-10,null,5]
	Explanation: [0,-10,5,null,-3,null,9] is also accepted:

Example 2:

	Input: nums = [1,3]
	Output: [3,1]
	Explanation: [1,null,3] and [3,1] are both height-balanced BSTs.

Constraints:

	- 1 <= nums.length <= 10^4
	- -10^4 <= nums[i] <= 10^4
	- nums is sorted in a strictly increasing order.
*/

func sortedArrayToBST(nums []int) *cs.TreeNode {
	if len(nums) == 0 {
		return nil
	} else if len(nums) == 1 {
		return &cs.TreeNode{Val: nums[0]}
	}

	middle := len(nums) / 2
	head := &cs.TreeNode{Val: nums[middle]}

	head.Left = sortedArrayToBST(nums[:middle])
	head.Right = sortedArrayToBST(nums[middle+1:])

	return head
}
