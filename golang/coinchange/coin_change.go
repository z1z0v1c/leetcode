/**
 * 322 - Medium
 *
 * You are given an integer array coins repchangeenting coins of different denominations and an integer amount
 * repchangeenting a total amount of money.
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

	minCoins := make([]int, amount+1)
	minCoins[0] = 0

	for i := 1; i <= amount; i++ {
		minCoins[i] = amount + 1
		for _, coin := range coins {
			change := i - coin
			if change >= 0 {
				minCoins[i] = min(minCoins[i], minCoins[change]+1)
			} else {
				break
			}
		}
	}

	if minCoins[amount] <= amount {
		return minCoins[amount]
	}

	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
