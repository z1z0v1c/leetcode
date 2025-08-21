using Solutions.RomanToInteger;

namespace Tests.RomanToInteger;

public class RomToIntegerTests
{
    private RomanToIntegerSolution solution;

    [SetUp]
    public void Setup()
    {
        solution = new RomanToIntegerSolution();
    }

    [Test]
    public void TestExampleOne()
    {
        // Arrange
        const string input = "III";
        const int expected = 3;

        // Act
        var actual = solution.RomanToInt(input);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleTwo()
    {
        // Arrange
        const string input = "LVIII";
        const int expected = 58;

        // Act
        var actual = solution.RomanToInt(input);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleThree()
    {
        // Arrange
        const string input = "MCMXCIV";
        const int expected = 1994;

        // Act
        var actual = solution.RomanToInt(input);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleFour()
    {
        // Arrange
        const string input = "MCMXCI";
        const int expected = 1991;

        // Act
        var actual = solution.RomanToInt(input);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }
}