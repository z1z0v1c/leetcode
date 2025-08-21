using Solutions.ValidParentheses;

namespace Tests.ValidParentheses;

public class ValidParenthesesTests
{
    private ValidParenthesesSolution solution;

    [SetUp]
    public void Setup()
    {
        solution = new ValidParenthesesSolution();
    }

    [Test]
    public void TestExampleOne()
    {
        // Arrange
        const string s = "()";

        // Act
        var actual = solution.IsValid(s);

        // Assert
        Assert.That(actual, Is.True);
    }

    [Test]
    public void TestExampleTwo()
    {
        // Arrange
        const string s = "()[]{}";

        // Act
        var actual = solution.IsValid(s);

        // Assert
        Assert.That(actual, Is.True);
    }

    [Test]
    public void TestExampleThree()
    {
        // Arrange
        const string s = "(]";

        // Act
        var actual = solution.IsValid(s);

        // Assert
        Assert.That(actual, Is.False);
    }

    [Test]
    public void TestExampleFour()
    {
        // Arrange
        const string s = "([])";

        // Act
        var actual = solution.IsValid(s);

        // Assert
        Assert.That(actual, Is.True);
    }
}