using Solutions.RemoveElement;

namespace Tests.RemoveElement;

public class RemoveElementTests
{
    private RemoveElementSolution solution;

    [SetUp]
    public void Setup()
    {
        solution = new RemoveElementSolution();
    }

    [Test]
    public void TestExampleOne()
    {
        // Arrange
        int[] nums = [3, 2, 2, 3];
        const int val = 3;
        int[] expectedNums = [2, 2, 0, 0];
        const int expectedLength = 2;

        // Act
        var actualLength = solution.RemoveElement(nums, val);

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
        int[] nums = [0, 1, 2, 2, 3, 0, 4, 2];
        const int val = 2;
        int[] expectedNums = [0, 1, 4, 0, 3, 0, 0, 0];
        const int expectedLength = 5;

        // Act
        var actualLength = solution.RemoveElement(nums, val);

        // Assert
        Assert.That(actualLength, Is.EqualTo(expectedLength));
        for (var i = 0; i < expectedLength; i++)
        {
            Assert.That(nums[i], Is.EqualTo(expectedNums[i]));
        }
    }
}