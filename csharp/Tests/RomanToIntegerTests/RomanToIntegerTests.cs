using Solutions.RomanToInteger;

namespace Tests.RomanToIntegerTests;

public class RomToIntegerTests
{
    private RomanToIntegerSolution romanToIntegerSolution;

    [SetUp]
    public void Setup() => romanToIntegerSolution = new();

    [Test]
    public void TestExampleOne()
    {
        string input = "III";

        int expected = 3;
        int actual = romanToIntegerSolution.RomanToInt(input);

        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleTwo()
    {
        string input = "LVIII";
        
        int expected = 58;
        int actual = romanToIntegerSolution.RomanToInt(input);

        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleThree()
    {
        string input = "MCMXCIV";
        
        int expected = 1994;
        int actual = romanToIntegerSolution.RomanToInt(input);

        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleFour()
    {
        string input = "MCMXCI";
        
        int expected = 1991;
        int actual = romanToIntegerSolution.RomanToInt(input);

        Assert.That(actual, Is.EqualTo(expected));
    }
}

