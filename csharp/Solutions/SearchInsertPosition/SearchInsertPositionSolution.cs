/// <summary>
/// 35 - Easy
/// 
/// Given a sorted array of distinct integers and a target value, return the index if the target is found.
/// If not, return the index where it would be if it were inserted in order.
/// 
/// You must write an algorithm with O(log n) runtime complexity.
/// 
/// Example 1:
///      Input: nums = [1,3,5,6], target = 5
///      Output: 2
/// 
/// Example 2:
///      Input: nums = [1,3,5,6], target = 2
///      Output: 1
/// 
/// Example 3:
///      Input: nums = [1,3,5,6], target = 7
///      Output: 4
///  
/// Constraints:
///      - 1 <= nums.length <= 104
///      - -104 <= nums[i] <= 104
///      - nums contains distinct values sorted in ascending order.
///      - -104 <= target <= 104
/// </summary>

namespace Solutions.SearchInsertPosition;

public class SearchInsertPositionSolution
{
    public int SearchInsert(int[] nums, int target)
    {
        var start = 0;
        var end = nums.Length - 1;

        while (start <= end)
        {
            var mid = (start + end) / 2;

            if (nums[mid] == target) return mid;

            if (target < nums[mid])
            {
                if (mid == 0) return 0;

                if (target > nums[mid - 1]) return mid;

                end = mid - 1;
            }
            else
            {
                if (mid == nums.Length - 1) return nums.Length;

                if (target < nums[mid + 1]) return mid + 1;

                start = mid + 1;
            }
        }

        return -1;
    }
}