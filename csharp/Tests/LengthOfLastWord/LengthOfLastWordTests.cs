using Solutions.LengthOfLastWord;

namespace Tests.LengthOfLastWord;

public class LengthOfLastWordTests
{
    private LengthOfLastWordSolution solution;

    [SetUp]
    public void Setup()
    {
        solution = new LengthOfLastWordSolution();
    }

    [Test]
    public void TestExampleOne()
    {
        // Arrange
        const string s = "Hello World";
        const int expected = 5;

        // Act
        var actual = solution.LengthOfLastWord(s);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleTwo()
    {
        // Arrange
        const string s = "   fly me   to   the moon  ";
        const int expected = 4;

        // Act
        var actual = solution.LengthOfLastWord(s);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleThree()
    {
        // Arrange
        const string s = "luffy is still joyboy";
        const int expected = 6;

        // Act
        var actual = solution.LengthOfLastWord(s);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }
}