using Solutions.RomanToInteger;

namespace Tests.RomanToInteger;

public class RomToIntegerTests
{
    private RomanToIntegerSolution solution;

    [SetUp]
    public void Setup() => solution = new();

    [Test]
    public void TestExampleOne()
    {
        string input = "III";

        int expected = 3;
        int actual = solution.RomanToInt(input);

        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleTwo()
    {
        string input = "LVIII";
        
        int expected = 58;
        int actual = solution.RomanToInt(input);

        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleThree()
    {
        string input = "MCMXCIV";
        
        int expected = 1994;
        int actual = solution.RomanToInt(input);

        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleFour()
    {
        string input = "MCMXCI";
        
        int expected = 1991;
        int actual = solution.RomanToInt(input);

        Assert.That(actual, Is.EqualTo(expected));
    }
}

