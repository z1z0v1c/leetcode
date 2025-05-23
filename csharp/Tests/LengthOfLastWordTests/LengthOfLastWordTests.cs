using Solutions.LengthOfLastWord;

namespace Tests.LengthOfLastWordTests;

public class LengthOfLastWordTests
{
    private LengthOfLastWordSolution lengthOfLastWord;

    [SetUp]
    public void Setup() => lengthOfLastWord = new();

    [Test]
    public void TestExampleOne()
    {
        string s = "Hello World";

        int expected = 5;
        int actual = lengthOfLastWord.LengthOfLastWord(s);
        
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleTwo()
    {
        string s = "   fly me   to   the moon  ";

        int expected = 4;
        int actual = lengthOfLastWord.LengthOfLastWord(s);
        
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleThree()
    {
        string s = "luffy is still joyboy";

        int expected = 6;
        int actual = lengthOfLastWord.LengthOfLastWord(s);
        
        Assert.That(actual, Is.EqualTo(expected));
    }
}

