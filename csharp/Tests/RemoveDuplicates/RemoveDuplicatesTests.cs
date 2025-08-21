using Solutions.RemoveDuplicates;

namespace Tests.RemoveDuplicates;

public class RemoveDuplicatesTests
{
    private RemoveDuplicatesSolution solution;

    [SetUp]
    public void Setup()
    {
        solution = new RemoveDuplicatesSolution();
    }

    [Test]
    public void TestExampleOne()
    {
        // Arrange
        int[] nums = [1, 1, 2];
        int[] expectedNums = [1, 2, 0];
        const int expectedLength = 2;

        // Act
        var actualLength = solution.RemoveDuplicates(nums);

        // Assert
        Assert.That(actualLength, Is.EqualTo(expectedLength));
        for (var i = 0; i < expectedLength; i++)
        {
            Assert.That(nums[i], Is.EqualTo(expectedNums[i]));
        }
    }

    [Test]
    public void TestExampleTwo()
    {
        // Arrange
        int[] nums = [0, 0, 1, 1, 1, 2, 2, 3, 3, 4];
        int[] expectedNums = [0, 1, 2, 3, 4, 0, 0, 0, 0, 0];
        const int expectedLength = 5;

        // Act
        var actualLength = solution.RemoveDuplicates(nums);

        // Assert
        Assert.That(actualLength, Is.EqualTo(expectedLength));
        for (var i = 0; i < expectedLength; i++)
        {
            Assert.That(nums[i], Is.EqualTo(expectedNums[i]));
        }
    }
}