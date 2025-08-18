using Solutions.RemoveElement;

namespace Tests.RemoveElement;

public class RemoveElementTests
{
    private RemoveElementSolution solution;

    [SetUp]
    public void Setup() => solution = new();

    [Test]
    public void TestExampleOne()
    {
        int[] nums = [3, 2, 2, 3];
        int val = 3;

        int[] expectedNums = [2, 2, 0, 0];
        int expectedLength = 2;
        int actualLength = solution.RemoveElement(nums, val);

        Assert.That(actualLength, Is.EqualTo(expectedLength));
        for (int i = 0; i < expectedLength; i++)
        {
            Assert.That(nums[i], Is.EqualTo(expectedNums[i]));
        }
    }

    [Test]
    public void TestExampleTwo()
    {
        int[] nums = [0, 1, 2, 2, 3, 0, 4, 2];
        int val = 2;

        int[] expectedNums = [0, 1, 4, 0, 3, 0, 0, 0];
        int expectedLength = 5;
        int actualLength = solution.RemoveElement(nums, val);

        Assert.That(actualLength, Is.EqualTo(expectedLength));
        for (int i = 0; i < expectedLength; i++)
        {
            Assert.That(nums[i], Is.EqualTo(expectedNums[i]));
        }
    }
}

