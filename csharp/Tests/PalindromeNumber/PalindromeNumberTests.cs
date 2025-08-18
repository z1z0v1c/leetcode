using Solutions.PalindromeNumber;

namespace Tests.PalindromeNumber;

public class PalindromeNumberTests
{
    private PalindromeNumberSolution solution;

    [SetUp]
    public void Setup() => solution = new();

    [Test]
    public void TestExampleOne()
    {
        int input = 121;
        Assert.That(solution.IsPalindrome(input), Is.True);
    }

    [Test]
    public void TestExampleTwo()
    {
        int input = -121;
        Assert.That(solution.IsPalindrome(input), Is.False);
    }

    [Test]
    public void TestExampleThree()
    {
        int input = 10;
        Assert.That(solution.IsPalindrome(input), Is.False);
    }
}

