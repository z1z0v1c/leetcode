using Solutions.PalindromeNumber;

namespace Tests.PalindromeNumberTests;

public class PalindromeNumberTests
{
    private PalindromeNumber palindromeNumber;

    [SetUp]
    public void Setup() => palindromeNumber = new();

    [Test]
    public void TestExampleOne()
    {
        int input = 121;
        Assert.That(palindromeNumber.IsPalindrome(input), Is.True);
    }

    [Test]
    public void TestExampleTwo()
    {
        int input = -121;
        Assert.That(palindromeNumber.IsPalindrome(input), Is.False);
    }

    [Test]
    public void TestExampleThree()
    {
        int input = 10;
        Assert.That(palindromeNumber.IsPalindrome(input), Is.False);
    }
}