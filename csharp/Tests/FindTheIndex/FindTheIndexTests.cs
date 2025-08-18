using Solutions.FindTheIndex;

namespace Tests.FindTheIndex;

public class FindTheIndexTests
{
    private FindTheIndexSolution solution;

    [SetUp]
    public void Setup() => solution = new();

    [Test]
    public void TestExampleOne()
    {
        string haystack = "sadbutsad";
        string needle = "sad";

        int expected = 0;
        int actual = solution.StrStr(haystack, needle);
        
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleTwo()
    {
        string haystack = "leetcode";
        string needle = "leeto";

        int expected = -1;
        int actual = solution.StrStr(haystack, needle);
        
        Assert.That(actual, Is.EqualTo(expected));
    }
}

