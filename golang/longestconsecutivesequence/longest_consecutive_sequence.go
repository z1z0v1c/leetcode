/**
 * 128 - Medium
 *
 * Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.
 * You must write an algorithm that runs in O(n) time.
 *
 * Example 1:
 *
 * 	Input: nums = [100,4,200,1,3,2]
 * 	Output: 4
 * 	Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
 *
 * Example 2:
 *
 * 	Input: nums = [0,3,7,2,5,8,4,6,0,1]
 * 	Output: 9
 *
 * Constraints:
 *
 * 	- 0 <= nums.length <= 10^5
 * 	- -10^9 <= nums[i] <= 10^9
 */
package longestconsecutivesequence

func longestConsecutive(nums []int) int {
	sequences := make(map[int]int, len(nums))
	longest := 0

	for _, num := range nums {
		if _, ok := sequences[num]; !ok {
			// Get neighbor sequences
			left := sequences[num-1]
			right := sequences[num+1]

			new := left + right + 1

			// Update only boundaries
			sequences[num-left] = new
			sequences[num+right] = new

			if new > longest {
				longest = new
			}
		}
	}

	return longest
}
