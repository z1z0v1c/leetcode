/**
 * 322 - Medium
 *
 * You are given an integer array coins representing coins of different denominations and an integer amount
 * representing a total amount of money.
 *
 * Return the fewest number of coins that you need to make up that amount. If that amount of money cannot be made
 * up by any combination of the coins, return -1.
 *
 * You may assume that you have an infinite number of each kind of coin.
 *
 *
 * Example 1:
 *
 * 	Input: coins = [1,2,5], amount = 11
 * 	Output: 3
 * 	Explanation: 11 = 5 + 5 + 1
 *
 * Example 2:
 *
 * 	Input: coins = [2], amount = 3
 * 	Output: -1
 *
 * Example 3:
 *
 * 	Input: coins = [1], amount = 0
 * 	Output: 0
 *
 * Constraints:
 *
 * 	- 1 <= coins.length <= 12
 * 	- 1 <= coins[i] <= 2^31 - 1
 * 	- 0 <= amount <= 10^4
 */
package coinchange

import "sort"

func coinChange(coins []int, amount int) int {
	sort.Ints(coins)

	mem := make([]int, amount+1)
	mem[0] = 0

	for i := 1; i <= amount; i++ {
		mem[i] = -1
	}

	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			res := i - coins[j]
			if res >= 0 {
				if (mem[i] == -1 || mem[i] > mem[res] + 1) && mem[res] != -1 {
					mem[i] = mem[res] + 1
				}
			} else {
				break
			}
		}
	}

	return mem[amount]
}
