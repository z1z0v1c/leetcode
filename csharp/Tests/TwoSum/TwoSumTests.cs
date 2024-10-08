using Solutions.TwoSum;

namespace Tests.TwoSumTests;

public class TwoSumTests
{
    private TwoSumSolution twoSum;

    [SetUp]
    public void Setup()
    {
        twoSum = new TwoSumSolution();
    }

    [Test]
    public void TestExampleOne()
    {
        int[] nums = [2,7,11,15];
        int target = 9;

        int[] actual = twoSum.TwoSum(nums, target);
        int[] expected = [0,1];

        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleTwo()
    {
        int[] nums = [3,2,4];
        int target = 6;

        int[] actual = twoSum.TwoSum(nums, target);
        int[] expected = [1,2];

        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleThree()
    {
        int[] nums = [3,3];
        int target = 6;

        int[] actual = twoSum.TwoSum(nums, target);
        int[] expected = [0, 1];

        Assert.That(actual, Is.EqualTo(expected));
    }
}
