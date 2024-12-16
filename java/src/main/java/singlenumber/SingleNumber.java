/**
 * 136 - Easy
 * <p>
 * Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
 * You must implement a solution with a linear runtime complexity and use only constant extra space.
 * <p>
 * Example 1:
 *      Input: nums = [2,2,1]
 *      Output: 1
 * <p>
 * Example 2:
 *      Input: nums = [4,1,2,1,2]
 *      Output: 4
 * <p>
 * Example 3:
 *      Input: nums = [1]
 *      Output: 1
 * <p>
 * Constraints:
 *      - 1 <= nums.length <= 3 * 10^4
 *      - -3 * 10^4 <= nums[i] <= 3 * 10^4
 *      - Each element in the array appears twice except for one element which appears only once.
 */
package singlenumber;

public class SingleNumber {
    public int singleNumber(int[] nums) {
        int result = 0;

        for (int num : nums) {
            result ^= num;
        }

        return result;
    }
}

