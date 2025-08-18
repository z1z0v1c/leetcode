/// <summary>
/// 20 - Easy
/// 
/// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', 
/// determine if the input string is valid.
/// 
/// An input string is valid if:
///      - Open brackets must be closed by the same type of brackets.
///      - Open brackets must be closed in the correct order.
///      - Every close bracket has a corresponding open bracket of the same type.
///  
/// Example 1:
///      Input: s = "()"
///      Output: true
/// 
/// Example 2:
///      Input: s = "()[]{}"
///      Output: true
/// 
/// Example 3:
///      Input: s = "(]"
///      Output: false
/// 
/// Example 4:
///      Input: s = "([])"
///      Output: true
/// 
/// Constraints:
///      - 1 <= s.length <= 104
///      - s consists of parentheses only '()[]{}'.
/// </summary>
namespace Solutions.ValidParentheses;

public class ValidParenthesesSolution
{
    public bool IsValid(string s)
    {
        Stack<char> stack = new();

        for (int i = 0; i < s.Length; i++)
        {
            if (s[i] == '(' || s[i] == '{' || s[i] == '[')
            {
                stack.Push(s[i]);
            }
            else if (s[i] == ')' || s[i] == '}' || s[i] == ']')
            {
                if (stack.Count == 0)
                {
                    return false;
                }

                char c = stack.Peek();

                if (
                    (c == '(' && s[i] == ')') ||
                    (c == '{' && s[i] == '}') ||
                    (c == '[' && s[i] == ']'))
                {
                    stack.Pop();
                }
                else
                {
                    return false;
                }
            }
        }

        if (stack.Count != 0)
        {
            return false;
        }

        return true;
    }
}

