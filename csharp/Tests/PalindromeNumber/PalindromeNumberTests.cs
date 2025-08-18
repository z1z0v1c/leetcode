using Solutions.PalindromeNumber;

namespace Tests.PalindromeNumber;

public class PalindromeNumberTests
{
    private Solutions.PalindromeNumber.PalindromeNumberSolution palindromeNumberSolution;

    [SetUp]
    public void Setup() => palindromeNumberSolution = new();

    [Test]
    public void TestExampleOne()
    {
        int input = 121;
        Assert.That(palindromeNumberSolution.IsPalindrome(input), Is.True);
    }

    [Test]
    public void TestExampleTwo()
    {
        int input = -121;
        Assert.That(palindromeNumberSolution.IsPalindrome(input), Is.False);
    }

    [Test]
    public void TestExampleThree()
    {
        int input = 10;
        Assert.That(palindromeNumberSolution.IsPalindrome(input), Is.False);
    }
}

