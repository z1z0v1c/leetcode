/**
 * 202 - Easy
 * <p>
 * Write an algorithm to determine if a number n is happy.
 * <p>
 * A happy number is a number defined by the following process:
 *      Starting with any positive integer, replace the number by the sum of the squares of its digits.
 *      Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
 *      Those numbers for which this process ends in 1 are happy.
 *      Return true if n is a happy number, and false if not.
 * <p>
 * Example 1:
 *      Input: n = 19
 *      Output: true
 *      Explanation:
 *          1^2 + 9^2 = 82
 *          8^2 + 2^2 = 68
 *          6^2 + 8^2 = 100
 *          1^2 + 0^2 + 0^2 = 1
 * <p>
 * Example 2:
 *      Input: n = 2
 *      Output: false
 * <p>
 * Constraints:
 *      - 1 <= n <= 2^31 - 1
 */
package happynumber;

import java.util.HashSet;
import java.util.Set;

public class HappyNumber {
    public boolean isHappy(int n) {
        if (n <= 1) {
            return true;
        }

        Set<Integer> checked = new HashSet<>();
        while (n > 1) {
            if (checked.contains(n)) {
                return false;
            }
            checked.add(n);

            int mod;
            int sum = 0;
            while (n > 0) {
                mod = n % 10;
                sum += (int) Math.pow(mod, 2);
                n /= 10;
            }

            if (sum == 1) {
                return true;
            }
            n = sum;
        }

        return false;
    }
}

