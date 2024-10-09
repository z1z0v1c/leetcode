using Solutions.RemoveDuplicates;

namespace Tests.RemoveDuplicatesTests;

public class RemoveDuplicatesTests
{
    private RemoveDuplicatesSolution removeDuplicates;

    [SetUp]
    public void Setup() => removeDuplicates = new();

    [Test]
    public void TestExampleOne()
    {
        int[] nums = [1, 1, 2];
        int[] expectedNums = [1, 2, 0];

        int expectedLength = 2;
        int actualLength = removeDuplicates.RemoveDuplicates(nums);

        Assert.That(actualLength, Is.EqualTo(expectedLength));
        for (int i = 0; i < expectedLength; i++)
        {
            Assert.That(nums[i], Is.EqualTo(expectedNums[i]));
        }
    }

    [Test]
    public void TestExampleTwo()
    {
        int[] nums = [0, 0, 1, 1, 1, 2, 2, 3, 3, 4];
        int[] expectedNums = [0, 1, 2, 3, 4, 0, 0, 0, 0, 0];

        int expectedLength = 5;
        int actualLength = removeDuplicates.RemoveDuplicates(nums);

        Assert.That(actualLength, Is.EqualTo(expectedLength));
        for (int i = 0; i < expectedLength; i++)
        {
            Assert.That(nums[i], Is.EqualTo(expectedNums[i]));
        }
    }
}
