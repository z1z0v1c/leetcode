using Solutions.CountHillsAndValleys;

namespace Tests.CountHillsAndValleys;

public class CountHillsAndValleysTests
{
    private CountHillsAndValleysSolution solution;

    [SetUp]
    public void Setup()
    {
        solution = new CountHillsAndValleysSolution();
    }

    [Test]
    public void TestExampleOne()
    {
        // Arrange
        var nums = new[] { 2, 4, 1, 1, 6, 5 };
        const int expected = 3;

        // Act
        var actual = solution.CountHillValley(nums);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleTwo()
    {
        // Arrange
        var nums = new[] { 6, 6, 5, 5, 4, 1 };
        const int expected = 0;

        // Act
        var actual = solution.CountHillValley(nums);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleThree()
    {
        // Arrange
        var nums = new[]
        {
            57, 57, 57, 57, 57, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 85, 85, 85, 86, 86,
            86
        };
        const int expected = 2;

        // Act
        var actual = solution.CountHillValley(nums);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }
}