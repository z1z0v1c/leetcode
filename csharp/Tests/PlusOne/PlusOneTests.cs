using Solutions.PlusOne;

namespace Tests.PlusOne;

public class PlusOneTests
{
    private PlusOneSolution solution;

    [SetUp]
    public void Setup()
    {
        solution = new PlusOneSolution();
    }

    [Test]
    public void TestExampleOne()
    {
        // Arrange
        int[] digits = [1, 2, 3];
        int[] expected = [1, 2, 4];

        // Act
        var actual = solution.PlusOne(digits);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleTwo()
    {
        // Arrange
        int[] digits = [4, 3, 2, 1];
        int[] expected = [4, 3, 2, 2];

        // Act
        var actual = solution.PlusOne(digits);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleThree()
    {
        // Arrange
        int[] digits = [9];
        int[] expected = [1, 0];

        // Act
        var actual = solution.PlusOne(digits);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }
}