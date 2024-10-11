package mergesortedarrays;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class MergeSortedArraysTest {
    private MergeSortedArrays mergeSortedArrays;

    @BeforeEach
    void setUp() {
        mergeSortedArrays = new MergeSortedArrays();
    }

    @Test
    void testExampleOne() {
        int[] nums1 = {1};
        int m = 1;
        int[] nums2 = null;
        int n = 0;

        int[] expected = {1};
        mergeSortedArrays.merge(nums1, m, nums2, n);

        for (int i = 0; i < nums1.length; i++) {
            assertEquals(expected[i], nums1[i], "Elements should be equal");
        }
    }
}
