using Solutions.PeakIndexInAnArray;

namespace Tests.PeakIndexInAnArray;

public class PeakIndexInAnArrayTests
{
    private PeakIndexInAnArraySolution solution;

    [SetUp]
    public void Setup() => solution = new();

    [Test]
    public void TestExampleOne()
    {
        // Arrange
        var arr = new[] { 0, 1, 0 };
        var expected = 1;

        // Act
        var actual = solution.PeakIndexInMountainArray(arr);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleTwo()
    {
        // Arrange
        var arr = new[] { 0, 2, 1, 0 };
        var expected = 1;

        // Act
        var actual = solution.PeakIndexInMountainArray(arr);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleThree()
    {
        // Arrange
        var arr = new[] { 0, 10, 5, 2 };
        var expected = 1;

        // Act
        var actual = solution.PeakIndexInMountainArray(arr);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleFour()
    {
        // Arrange
        var arr = new[] { 40, 48, 61, 75, 100, 99, 98, 39, 30, 10 };
        var expected = 4;

        // Act
        var actual = solution.PeakIndexInMountainArray(arr);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleFive()
    {
        // Arrange
        var arr = new[] { 3, 4, 5, 1 };
        var expected = 2;

        // Act
        var actual = solution.PeakIndexInMountainArray(arr);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleSix()
    {
        // Arrange
        var arr = new[] { 18, 29, 38, 59, 98, 100, 99, 98, 90 };
        var expected = 5;

        // Act
        var actual = solution.PeakIndexInMountainArray(arr);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }
}