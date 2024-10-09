using Solutions.ValidParentheses;

namespace Tests.ValidParenthesesTests;

public class ValidParenthesesTests
{
    private ValidParentheses validParentheses;

    [SetUp]
    public void Setup() => validParentheses = new();

    [Test]
    public void TestExampleOne()
    {
        string s = "()";

        Assert.That(validParentheses.IsValid(s), Is.True);
    }

    [Test]
    public void TestExampleTwo()
    {
        string s = "()[]{}";

        Assert.That(validParentheses.IsValid(s), Is.True);
    }

    [Test]
    public void TestExampleThree()
    {
        string s = "(]";

        Assert.That(validParentheses.IsValid(s), Is.False);
    }

    [Test]
    public void TestExampleFour()
    {
        string s = "([])";

        Assert.That(validParentheses.IsValid(s), Is.True);
    }
}
