using Solutions.IsomorphicStrings;

namespace Tests.IsomorphicStrings;

public class IsomorphicStringsTests
{
    private IsomorphicStringsSolution solution;

    [SetUp]
    public void SetUp()
    {
        solution = new IsomorphicStringsSolution();
    }

    [Test]
    public void TestExampleOne()
    {
        // Arrange
        const string s = "egg";
        const string t = "add";

        // Act
        var result = solution.IsIsomorphic(s, t);

        // Assert
        Assert.That(result, Is.True);
    }

    [Test]
    public void TestExampleTwo()
    {
        // Arrange
        const string s = "foo";
        const string t = "bar";

        // Act
        var result = solution.IsIsomorphic(s, t);

        // Assert
        Assert.That(result, Is.False);
    }

    [Test]
    public void TestExampleThree()
    {
        // Arrange
        const string s = "paper";
        const string t = "title";

        // Act
        var result = solution.IsIsomorphic(s, t);

        // Assert
        Assert.That(result, Is.True);
    }

    [Test]
    public void TestExampleFour()
    {
        // Arrange
        const string s = "bbbaaaba";
        const string t = "aaabbbba";

        // Act
        var result = solution.IsIsomorphic(s, t);

        // Assert
        Assert.That(result, Is.False);
    }

    [Test]
    public void TestExampleFive()
    {
        // Arrange
        const string s = "badc";
        const string t = "baba";

        // Act
        var result = solution.IsIsomorphic(s, t);

        // Assert
        Assert.That(result, Is.False);
    }
}