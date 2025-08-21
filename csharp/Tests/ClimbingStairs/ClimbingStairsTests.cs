using Solutions.ClimbingStairs;

namespace Tests.ClimbingStairs;

public class ClimbingStairsTests
{
    private ClimbingStairsSolution solution;

    [SetUp]
    public void Setup()
    {
        solution = new ClimbingStairsSolution();
    }

    [Test]
    public void TestExampleOne()
    {
        // Arrange
        const int n = 2;
        const int expected = 2;

        // Act
        var actual = solution.ClimbStairs(n);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleTwo()
    {
        // Arrange
        const int n = 3;
        const int expected = 3;

        // Act
        var actual = solution.ClimbStairs(n);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }
}