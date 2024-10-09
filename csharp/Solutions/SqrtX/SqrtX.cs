namespace Solutions.SqrtX;

/// <summary>
/// 69 - Easy
/// 
/// Given a non-negative integer x, return the square root of x rounded down to the nearest integer. 
/// The returned integer should be non-negative as well.
/// 
/// You must not use any built-in exponent function or operator.
/// For example, do not use pow(x, 0.5) in c++ or x ** 0.5 in python.
///  
/// Example 1:
///      Input: x = 4
///      Output: 2
///      Explanation: The square root of 4 is 2, so we return 2.
/// 
/// Example 2:
///      Input: x = 8
///      Output: 2
///      Explanation: 
///          The square root of 8 is 2.82842..., and since we round it 
///          down to the nearest integer, 2 is returned.
///  
/// Constraints:
///      - 0 <= x <= 231 - 1
/// </summary>

public class SqrtX
{
    public int MySqrt(int x) {
        if (x == 0 || x == 1)
        {
            return x;
        }   

        int start = 0;
        int end = x;

        while (start <= end)
        {
            int mid = (start + end) / 2;
            int square = mid * mid;

            if (square == x)
            {
                return mid;
            }
            else if (mid <= x / mid)
            {
                if ((mid + 1) > x / (mid + 1))
                {
                    return mid;
                }

                start = mid;
            }
            else
            {
                if ((mid - 1) < x / (mid - 1))
                {
                    return mid - 1;
                }

                end = mid;
            }
        }

        return end;
    }
}
