using PalindromeNumber;

namespace PalindromeNumberTests;

public class PalindromeNumberTests
{
    private PalindromeNumberSolution palindromeNumber;

    [SetUp]
    public void Setup()
    {
        palindromeNumber = new PalindromeNumberSolution();
    }

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