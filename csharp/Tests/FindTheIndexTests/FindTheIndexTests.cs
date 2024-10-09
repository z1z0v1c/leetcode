using Solutions.FindTheIndex;

namespace Tests.FindTheIndexTests;

public class FindTheIndexTests
{
    private FindTheIndex findTheIndex;

    [SetUp]
    public void Setup() => findTheIndex = new();

    [Test]
    public void TestExampleOne()
    {
        string haystack = "sadbutsad";
        string needle = "sad";

        int expected = 0;
        int actual = findTheIndex.StrStr(haystack, needle);
        
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleTwo()
    {
        string haystack = "leetcode";
        string needle = "leeto";

        int expected = -1;
        int actual = findTheIndex.StrStr(haystack, needle);
        
        Assert.That(actual, Is.EqualTo(expected));
    }
}
