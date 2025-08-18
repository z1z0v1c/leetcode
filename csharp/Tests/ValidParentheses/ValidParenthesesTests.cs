using Solutions.ValidParentheses;

namespace Tests.ValidParentheses;

public class ValidParenthesesTests
{
    private ValidParenthesesSolution solution;

    [SetUp]
    public void Setup() => solution = new();

    [Test]
    public void TestExampleOne()
    {
        string s = "()";

        Assert.That(solution.IsValid(s), Is.True);
    }

    [Test]
    public void TestExampleTwo()
    {
        string s = "()[]{}";

        Assert.That(solution.IsValid(s), Is.True);
    }

    [Test]
    public void TestExampleThree()
    {
        string s = "(]";

        Assert.That(solution.IsValid(s), Is.False);
    }

    [Test]
    public void TestExampleFour()
    {
        string s = "([])";

        Assert.That(solution.IsValid(s), Is.True);
    }
}

