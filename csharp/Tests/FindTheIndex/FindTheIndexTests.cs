using Solutions.FindTheIndex;

namespace Tests.FindTheIndex;

public class FindTheIndexTests
{
    private FindTheIndexSolution solution;

    [SetUp]
    public void Setup()
    {
        solution = new FindTheIndexSolution();
    }

    [Test]
    public void TestExampleOne()
    {
        // Arrange
        const string haystack = "sadbutsad";
        const string needle = "sad";
        const int expected = 0;

        // Act
        var actual = solution.StrStr(haystack, needle);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleTwo()
    {
        // Arrange
        const string haystack = "leetcode";
        const string needle = "leeto";
        const int expected = -1;

        // Act
        var actual = solution.StrStr(haystack, needle);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }
}