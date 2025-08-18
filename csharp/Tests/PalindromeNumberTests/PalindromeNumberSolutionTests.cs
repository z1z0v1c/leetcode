using Solutions.PalindromeNumber;

namespace Tests.PalindromeNumberTests;

public class PalindromeNumberSolutionTests
{
    private PalindromeNumberSolution palindromeNumberSolution;

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

