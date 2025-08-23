using Solutions.SummaryRanges;

namespace Tests.SummaryRanges;

public class SummaryRangesTests
{
    private SummaryRangesSolution solution;

    [SetUp]
    public void Setup()
    {
        solution = new SummaryRangesSolution();
    }

    [Test]
    public void TestExampleOne()
    {
        // Arrange
        var nums = new[] { 0, 1, 2, 4, 5, 7 };
        var expected = new[] { "0->2", "4->5", "7" };

        // Act
        var actual = solution.SummaryRanges(nums);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleTwo()
    {
        // Arrange
        var nums = new[] { 0, 2, 3, 4, 6, 8, 9 };
        var expected = new[] { "0", "2->4", "6", "8->9" };

        // Act
        var actual = solution.SummaryRanges(nums);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleThree()
    {
        // Arrange
        int[] nums = [];
        string[] expected = [];

        // Act
        var actual = solution.SummaryRanges(nums);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleFour()
    {
        // Arrange
        var nums = new[] { -2147483648, 0, 2, 3, 4, 6, 8, 9 };
        var expected = new[] { "-2147483648", "0", "2->4", "6", "8->9" };

        // Act
        var actual = solution.SummaryRanges(nums);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }
}