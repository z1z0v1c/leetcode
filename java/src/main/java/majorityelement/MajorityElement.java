/**
 * Given an array nums of size n, return the majority element.
 * <p>
 * The majority element is the element that appears more than ⌊n / 2⌋ times.
 * You may assume that the majority element always exists in the array.
 * <p>
 * Example 1:
 *      Input: nums = [3,2,3]
 *      Output: 3
 * <p>
 * Example 2:
 *      Input: nums = [2,2,1,1,1,2,2]
 *      Output: 2
 * <p>
 * Constraints:
 *      - n == nums.length
 *      - 1 <= n <= 5 * 10^4
 *      - -10^9 <= nums[i] <= 10^9
 * <p>
 * Follow-up: Could you solve the problem in linear time and in O(1) space?
 */
package majorityelement;

import java.util.Arrays;

public class MajorityElement {
    public int majorityElement(int[] nums) {
        Arrays.sort(nums);

        int num = nums[0];
        int count = 1;

        for (int i = 1; i < nums.length; i ++) {
            if (nums[i] == num) {
                count++;
            } else {
                num = nums[i];
                count = 1;
            }

            if (count > nums.length / 2) {
                return num;
            }
        }

        return num;
    }
}

