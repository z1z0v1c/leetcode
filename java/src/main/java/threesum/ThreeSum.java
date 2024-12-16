/**
 * 15 - Medium
 * <p>
 * Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that
 * i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
 * <p>
 * Notice that the solution set must not contain duplicate triplets.
 * <p>
 * Example 1:
 *      Input: nums = [-1,0,1,2,-1,-4]
 *      Output: [[-1,-1,2],[-1,0,1]]
 *      Explanation:
 *          nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
 *          nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
 *          nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
 *          The distinct triplets are [-1,0,1] and [-1,-1,2].
 *          Notice that the order of the output and the order of the triplets does not matter.
 * <p>
 * Example 2:
 *      Input: nums = [0,1,1]
 *      Output: []
 *      Explanation: The only possible triplet does not sum up to 0.
 * <p>
 * Example 3:
 *      Input: nums = [0,0,0]
 *      Output: [[0,0,0]]
 *      Explanation: The only possible triplet sums up to 0.
 * <p>
 * Constraints:
 *      - 3 <= nums.length <= 3000
 *      - -105 <= nums[i] <= 105
 */
package threesum;

import java.util.*;

public class ThreeSum {
    public List<List<Integer>> threeSum(int[] nums) {
        Arrays.sort(nums);
        Set<List<Integer>> threes = new HashSet<>();

        for (int i = 0; i < nums.length - 2; i++) {
            if (nums[i] > 0) {
                break;
            }
            for (int j = i + 1; j < nums.length - 1; j++) {
                if (nums[i] + nums[j] > 0) {
                    break;
                }
                int index = Arrays.binarySearch(nums, j + 1, nums.length, -nums[i] - nums[j]);
                if (index > 0) {
                    List<Integer> three = new ArrayList<>();
                    three.add(nums[i]);
                    three.add(nums[j]);
                    three.add(nums[index]);

                    threes.add(three);
                }
            }
        }

        return threes.stream().toList();
    }
}

