/// <summary>
/// 14 - Easy
/// 
/// Write a function to find the longest common prefix string amongst an array of strings.
/// If there is no common prefix, return an empty string "".
/// 
/// Example 1:
///      Input: strs = ["flower","flow","flight"]
///      Output: "fl"
/// 
/// Example 2:
///      Input: strs = ["dog","racecar","car"]
///      Output: ""
///      Explanation: There is no common prefix among the input strings.
///  
/// Constraints:
///      - 1 <= strs.length <= 200
///      - 0 <= strs[i].length <= 200
///      - strs[i] consists of only lowercase English letters.
/// </summary>
namespace Solutions.LongestCommonPrefix;

public class LongestCommonPrefixSolution
{
    public string LongestCommonPrefix(string[] strs)
    {
        string prefix = "";

        for (int i = 0; i < strs[0].Length; i++)
        {
            for (int j = 0; j < strs.Length; j ++)
            {
                if (i >= strs[j].Length || strs[0][i] != strs[j][i])
                {
                    return prefix;
                }                    
            }
                prefix += strs[0][i];
        }

        return prefix;
    }
}

