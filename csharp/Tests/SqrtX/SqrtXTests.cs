using Solutions.SqrtX;

namespace Tests.SqrtX;

public class SqrtXTests
{
    private SqrtXSolution solution;

    [SetUp]
    public void Setup()
    {
        solution = new SqrtXSolution();
    }

    [Test]
    public void TestExampleOne()
    {
        // Arrange
        const int x = 4;
        const int expected = 2;

        // Act
        var actual = solution.MySqrt(x);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleTwo()
    {
        // Arrange
        const int x = 8;
        const int expected = 2;

        // Act
        var actual = solution.MySqrt(x);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }
}