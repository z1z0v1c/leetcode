using System.Numerics;

namespace Solutions.AddBinary;

/// <summary>
/// 67 - Easy
/// 
/// Given two binary strings a and b, return their sum as a binary string.
/// 
/// Example 1:
///      Input: a = "11", b = "1"
///      Output: "100"
/// 
/// Example 2:
///      Input: a = "1010", b = "1011"
///      Output: "10101"
///  
/// Constraints:
///      - 1 <= a.length, b.length <= 104
///      - a and b consist only of '0' or '1' characters.
///      - Each string does not contain leading zeros except for the zero itself.
/// </summary>

public class AddBinarySolution {
    public string AddBinary(string a, string b) {
        // Convert binary strings to BigInteger
        BigInteger num1 = BinaryStringToBigInteger(a);
        BigInteger num2 = BinaryStringToBigInteger(b);

        // Compute the sum using BigInteger addition
        BigInteger sum = num1 + num2;

        // Convert the result to a binary string
        return ToBinaryString(sum);
    }

    private static BigInteger BinaryStringToBigInteger(string binary)
    {
        BigInteger result = 0;
        foreach (char bit in binary)
        {
            result <<= 1;
            if (bit == '1')
            {
                result |= 1;
            }
        }
        return result;
    }

    private static string ToBinaryString(BigInteger value)
    {
        if (value == 0) return "0";
        string binary = "";
        while (value > 0)
        {
            binary = (value % 2) + binary;
            value >>= 1;
        }
        return binary;
    }
}
