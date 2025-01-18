/**
 * 875 - Medium
 *
 * Koko loves to eat bananas. There are n piles of bananas, the ith pile has piles[i] bananas. The guards have gone and will come back in h hours.
 *
 * Koko can decide her bananas-per-hour eating speed of k. Each hour, she chooses some pile of bananas and eats k bananas from that pile.
 * If the pile has less than k bananas, she eats all of them instead and will not eat any more bananas during this hour.
 *
 * Koko likes to eat slowly but still wants to finish eating all the bananas before the guards return.
 *
 * Return the minimum integer k such that she can eat all the bananas within h hours.
 *
 *
 * Example 1:
 *
 * 	Input: piles = [3,6,7,11], h = 8
 * 	Output: 4
 *
 * Example 2:
 *
 * 	Input: piles = [30,11,23,4,20], h = 5
 * 	Output: 30
 *
 * Example 3:
 *
 * 	Input: piles = [30,11,23,4,20], h = 6
 * 	Output: 23
 *
 * Constraints:
 *
 * 	- 1 <= piles.length <= 10^4
 * 	- piles.length <= h <= 10^9
 * 	- 1 <= piles[i] <= 10^9
 */
package kokoeatingbananas

func minEatingSpeed(piles []int, h int) int {
	min, max := 1, max(piles)
	for min <= max {
		speed := (min + max) / 2
		if isEatable(piles, speed, h) {
			max = speed - 1
		} else {
			min = speed + 1
		}
	}

	return min
}

func max(piles []int) int {
	var max int 
	for _, pile := range piles {
		if pile > max {
			max = pile
		}
	}

	return max
}

func isEatable(piles []int, speed int, h int) bool {
	for _, pile := range piles {
		if pile % speed == 0 {
			h -= pile / speed
		} else {
			h -= pile / speed + 1
		}
	}

	return h >= 0
}
