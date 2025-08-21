using Solutions.TwoSum;

namespace Tests.TwoSum;

public class TwoSumTests
{
    private TwoSumSolution solution;

    [SetUp]
    public void Setup()
    {
        solution = new TwoSumSolution();
    }

    [Test]
    public void TestExampleOne()
    {
        // Arrange
        int[] nums = [2, 7, 11, 15];
        const int target = 9;
        int[] expected = [0, 1];

        // Act
        var actual = solution.TwoSum(nums, target);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleTwo()
    {
        // Arrange
        int[] nums = [3, 2, 4];
        const int target = 6;
        int[] expected = [1, 2];

        // Act
        var actual = solution.TwoSum(nums, target);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleThree()
    {
        // Arrange
        int[] nums = [3, 3];
        const int target = 6;
        int[] expected = [0, 1];

        // Act
        var actual = solution.TwoSum(nums, target);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }
}