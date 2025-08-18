using Solutions.LongestCommonPrefix;

namespace Tests.LongestCommonPrefix;

public class LittleLongestCommonPrefixTests
{
    private LongestCommonPrefixSolution longestCommonPrefix;

    [SetUp]
    public void Setup() => longestCommonPrefix = new();

    [Test]
    public void TestExampleOne()
    {
        string[] strs = ["flower","flow","flight"];

        string expected = "fl";
        string actual = longestCommonPrefix.LongestCommonPrefix(strs);
        
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleTwo()
    {
        string[] strs = ["dog","racecar","car"];

        string expected = "";
        string actual = longestCommonPrefix.LongestCommonPrefix(strs);
        
        Assert.That(actual, Is.EqualTo(expected));
    }
}

