package threesumclosest;

import java.util.*;

/**
 * 16 Medium
 * <p>
 * Given an integer array nums of length n and an integer target, find three integers in nums such that the sum is closest to target.
 * Return the sum of the three integers.
 * You may assume that each input would have exactly one solution.
 * <p>
 * Example 1:
 *      Input: nums = [-1,2,1,-4], target = 1
 *      Output: 2
 *      Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
 * <p>
 * Example 2:
 *      Input: nums = [0,0,0], target = 1
 *      Output: 0
 *      Explanation: The sum that is closest to the target is 0. (0 + 0 + 0 = 0).
 * <p>
 * Constraints:
 *      - 3 <= nums.length <= 500
 *      - -1000 <= nums[i] <= 1000
 *      - -104 <= target <= 104
 */

public class ThreeSumClosest {
    public int threeSumClosest(int[] nums, int target) {
        Arrays.sort(nums);

        int minDiff = Integer.MAX_VALUE;
        int min = 0;

        for (int i = 0; i < nums.length - 2; i++) {
            int left = i + 1;
            int right = nums.length - 1;


            while (left < right) {
                int sum = nums[i] + nums[left] + nums[right];

                if (sum == target) {
                    return sum;
                } else if (sum > target) {
                    right--;
                } else {
                    left++;
                }

                int currentDiff = Math.abs(sum - target);
                if (minDiff > currentDiff) {
                    minDiff = currentDiff;
                    min = sum;
                }
            }
        }

        return min;
    }
}
