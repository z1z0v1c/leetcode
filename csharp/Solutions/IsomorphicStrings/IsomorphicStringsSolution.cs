/// <sumary>
///
/// Given two strings s and t, determine if they are isomorphic. Two strings s and t are isomorphic if the characters in
/// s can be replaced to get t. All occurrences of a character must be replaced with another character while preserving
/// the order of characters. No two characters may map to the same character, but a character may map to itself.
/// 
/// Example 1:
///
///     Input: s = "egg", t = "add"
///     Output: true
///     Explanation: The strings s and t can be made identical by:
///         - Mapping 'e' to 'a'.
///         - Mapping 'g' to 'd'.
///
/// Example 2:
/// 
///     Input: s = "foo", t = "bar"
///     Output: false
///     Explanation: The strings s and t can not be made identical as 'o' needs to be mapped to both 'a' and 'r'.
/// 
/// Example 3:
/// 
///     Input: s = "paper", t = "title"
///     Output: true
/// 
/// Constraints:
/// 
///     - 1 <= s.length <= 5 * 10^4
///     - t.length == s.length
///     - s and t consist of any valid ascii character.
/// </sumary>
namespace Solutions.IsomorphicStrings;

public class IsomorphicStringsSolution
{
    public bool IsIsomorphic(string s, string t)
    {
        var sPairs = new Dictionary<char, char>();
        var tPairs = new Dictionary<char, char>();

        for (var i = 0; i < s.Length; i++)
        {
            if (!sPairs.ContainsKey(s[i]))
            {
                sPairs[s[i]] = t[i];
            }
            else
            {
                if (sPairs[s[i]] != t[i])
                {
                    return false;
                }
            }
            
            if (!tPairs.ContainsKey(t[i]))
            {
                tPairs[t[i]] = s[i];
            }
            else
            {
                if (tPairs[t[i]] != s[i])
                {
                    return false;
                }
            } 
        }
        
        return true;
    } 
}