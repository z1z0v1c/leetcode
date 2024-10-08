using Solutions.RomanToInteger;

namespace Tests.RomanToIntegerTests;

public class RomToIntegerTests
{
    private RomanToInteger romanToInteger;

    [SetUp]
    public void Setup()
    {
        romanToInteger = new RomanToInteger();
    }

    [Test]
    public void TestExampleOne()
    {
        string input = "III";

        int expected = 3;
        int actual = romanToInteger.RomanToInt(input);

        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleTwo()
    {
        string input = "LVIII";
        
        int expected = 58;
        int actual = romanToInteger.RomanToInt(input);

        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleThree()
    {
        string input = "MCMXCIV";
        
        int expected = 1994;
        int actual = romanToInteger.RomanToInt(input);

        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleFour()
    {
        string input = "MCMXCI";
        
        int expected = 1991;
        int actual = romanToInteger.RomanToInt(input);

        Assert.That(actual, Is.EqualTo(expected));
    }
}
