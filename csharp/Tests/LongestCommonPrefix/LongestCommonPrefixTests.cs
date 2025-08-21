using Solutions.LongestCommonPrefix;

namespace Tests.LongestCommonPrefix;

public class LittleLongestCommonPrefixTests
{
    private LongestCommonPrefixSolution solution;

    [SetUp]
    public void Setup()
    {
        solution = new LongestCommonPrefixSolution();
    }

    [Test]
    public void TestExampleOne()
    {
        // Arrange
        string[] strs = ["flower", "flow", "flight"];
        const string expected = "fl";

        // Act
        var actual = solution.LongestCommonPrefix(strs);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleTwo()
    {
        // Arrange
        string[] strs = ["dog", "racecar", "car"];
        const string expected = "";

        // Act
        var actual = solution.LongestCommonPrefix(strs);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }
}