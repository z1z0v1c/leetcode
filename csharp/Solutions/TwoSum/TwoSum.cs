namespace Solutions.TwoSum;

/// <summary>
/// 1 - Easy
/// 
/// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
/// You may assume that each input would have exactly one solution, and you may not use the same element twice.
/// You can return the answer in any order.
/// 
/// Example 1:
///      Input: nums = [2,7,11,15], target = 9
///      Output: [0,1]
///      Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
/// 
/// Example 2:
///      Input: nums = [3,2,4], target = 6
///      Output: [1,2]
///
/// Example 3:
///      Input: nums = [3,3], target = 6
///      Output: [0,1]
///
/// 
/// Constraints:
///      2 <= nums.length <= 104
///      -109 <= nums[i] <= 109
///      -109 <= target <= 109
/// 
/// Only one valid answer exists.
/// 
/// Follow-up: Can you come up with an algorithm that is less than O(n2) time complexity?
/// </summary>

public class TwoSumSolution
{
	public int[] TwoSum(int[] nums, int target)
	{
		// Store indexes
		Dictionary<int, List<int>> dict = PopulateDict(nums);

		Array.Sort(nums);

		int sum = 0;

		int left = 0;
		int right = nums.Length - 1;

		while (left <= right)
		{
			sum = nums[left] + nums[right];

			if (sum == target)
			{
				int first = dict.GetValueOrDefault(nums[left]).FirstOrDefault();
				int second = dict.GetValueOrDefault(nums[right]).LastOrDefault();
				return [first, second];
			}
			else if (sum > target)
			{
				right--;
			}
			else 
			{
				left++;
			}
		}

		return [-1, -1];
	}

	private Dictionary<int, List<int>> PopulateDict(int[] nums)
	{
		Dictionary<int, List<int>> dict = [];
		for (int i = 0; i < nums.Length; i++) {
			if (dict.ContainsKey(nums[i])) {
				dict.GetValueOrDefault(nums[i]).Add(i);
			}
			else
			{
				List<int> value = [];
				value.Add(i);
				dict.TryAdd(nums[i], value);
			}
			
		}

		return dict;
	}
}
