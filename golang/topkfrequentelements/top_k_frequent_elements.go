/**
 * 347 - Medium
 *
 * Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.
 *
 * Example 1:
 *
 * 	Input: nums = [1,1,1,2,2,3], k = 2
 * 	Output: [1,2]
 *
 * Example 2:
 *
 * 	Input: nums = [1], k = 1
 * 	Output: [1]
 *
 * Constraints:
 *
 * 	- 1 <= nums.length <= 10^5
 * 	- -10^4 <= nums[i] <= 10^4
 * 	- k is in the range [1, the number of unique elements in the array].
 * 	- It is guaranteed that the answer is unique.
 *
 * Follow up: Your algorithm's time complexity must be better than O(n log n), where n is the array's size.
 */
package topkfrequentelements

import "slices"

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)

	// Count frequencies
	for _, num := range nums {
		freq[num]++
	}

	elements := make([]int, 0, len(freq))
	for element := range freq {
		elements = append(elements, element)
	}

	// Sort elements by frequencies in descending order
	slices.SortFunc(elements, func(a, b int) int {
		if freq[a] > freq[b] {
			return -1
		}
		return 1
	})

	// Return top k elements
	return elements[:k]
}
