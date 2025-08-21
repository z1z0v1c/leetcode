using Solutions.PalindromeNumber;

namespace Tests.PalindromeNumber;

public class PalindromeNumberTests
{
    private PalindromeNumberSolution solution;

    [SetUp]
    public void Setup()
    {
        solution = new PalindromeNumberSolution();
    }

    [Test]
    public void TestExampleOne()
    {
        // Arrange
        const int input = 121;

        // Act
        var actual = solution.IsPalindrome(input);

        // Assert
        Assert.That(actual, Is.True);
    }

    [Test]
    public void TestExampleTwo()
    {
        // Arrange
        const int input = -121;

        // Act
        var actual = solution.IsPalindrome(input);

        // Assert
        Assert.That(actual, Is.False);
    }

    [Test]
    public void TestExampleThree()
    {
        // Arrange
        const int input = 10;

        // Act
        var actual = solution.IsPalindrome(input);

        // Assert
        Assert.That(actual, Is.False);
    }
}