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

    @Test
    void testExampleTwo() {
        int[] nums1 = {0};
        int m = 0;
        int[] nums2 = {1};
        int n = 1;

        int[] expected = {1};
        mergeSortedArrays.merge(nums1, m, nums2, n);

        for (int i = 0; i < nums1.length; i++) {
            assertEquals(expected[i], nums1[i], "Elements should be equal");
        }
    }

    @Test
    void testExampleThree() {
        int[] nums1 = {0, 0, 0};
        int m = 0;
        int[] nums2 = {3, 5, 9};
        int n = 3;

        int[] expected = {3, 5, 9};
        mergeSortedArrays.merge(nums1, m, nums2, n);

        for (int i = 0; i < nums1.length; i++) {
            assertEquals(expected[i], nums1[i], "Elements should be equal");
        }
    }

    @Test
    void testExampleFour() {
        int[] nums1 = {2, 0};
        int m = 1;
        int[] nums2 = {1};
        int n = 1;

        int[] expected = {1, 2, 2, 3, 5, 6};
        mergeSortedArrays.merge(nums1, m, nums2, n);

        for (int i = 0; i < nums1.length; i++) {
            assertEquals(expected[i], nums1[i], "Elements should be equal");
        }
    }

    @Test
    void testExampleFive() {
        int[] nums1 = {1, 0};
        int m = 1;
        int[] nums2 = {2};
        int n = 1;

        int[] expected = {1, 2, 2, 3, 5, 6};
        mergeSortedArrays.merge(nums1, m, nums2, n);

        for (int i = 0; i < nums1.length; i++) {
            assertEquals(expected[i], nums1[i], "Elements should be equal");
        }
    }
}
