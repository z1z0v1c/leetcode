using Solutions.FindThePeaks;

namespace Tests.FindThePeaks;

public class FindThePeaksTests
{
    private FindThePeaksSolution solution;

    [SetUp]
    public void Setup()
    {
        solution = new FindThePeaksSolution();
    }

    [Test]
    public void TestExampleOne()
    {
        // Arrange
        var mountain = new[] { 2, 4, 4 };
        var expected = new List<int>();

        // Act
        var actual = solution.FindPeaks(mountain);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleTwo()
    {
        // Arrange
        var mountain = new[] { 1, 4, 3, 8, 5 };
        var expected = new List<int> { 1, 3 };

        // Act
        var actual = solution.FindPeaks(mountain);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }
}