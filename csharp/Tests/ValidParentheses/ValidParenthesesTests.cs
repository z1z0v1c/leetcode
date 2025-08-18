using Solutions.ValidParentheses;

namespace Tests.ValidParentheses;

public class ValidParenthesesTests
{
    private Solutions.ValidParentheses.ValidParenthesesSolution validParenthesesSolution;

    [SetUp]
    public void Setup() => validParenthesesSolution = new();

    [Test]
    public void TestExampleOne()
    {
        string s = "()";

        Assert.That(validParenthesesSolution.IsValid(s), Is.True);
    }

    [Test]
    public void TestExampleTwo()
    {
        string s = "()[]{}";

        Assert.That(validParenthesesSolution.IsValid(s), Is.True);
    }

    [Test]
    public void TestExampleThree()
    {
        string s = "(]";

        Assert.That(validParenthesesSolution.IsValid(s), Is.False);
    }

    [Test]
    public void TestExampleFour()
    {
        string s = "([])";

        Assert.That(validParenthesesSolution.IsValid(s), Is.True);
    }
}

