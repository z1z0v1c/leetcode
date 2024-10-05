// Write a function that, given an array A of N integers,
// returns the smallest positive integer (greater than 0) that does not occur in A.

// For example, given A = [1, 3, 6, 4, 1, 2], the function should return 5.
// Given A = [1, 2, 3], the function should return 4.
// Given A = [−1, −3], the function should return 1.

// Write an efficient algorithm for the following assumptions:
//      N is an integer within the range [1..100,000];
//      each element of array A is an integer within the range [−1,000,000..1,000,000].

package smallestpositiveinteger;

import java.util.Arrays;

public class SmallestPositiveInteger {
    public int smallestPositiveInteger(int[] A) {
        Arrays.sort(A);

        int highest = 0;
        for (int num : A) {
            if (num < 1) {
                continue;
            }

            if ((num - highest) > 1) {
                return highest + 1;
            }

            highest = num;
        }

        return highest + 1;
    }
}
