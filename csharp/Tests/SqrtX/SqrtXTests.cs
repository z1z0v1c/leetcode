using Solutions.SqrtX;

namespace Tests.SqrtX;

public class SqrtXTests
{
    private SqrtXSolution solution;

    [SetUp]
    public void Setup() => solution = new();

    [Test]
    public void TestExampleOne()
    {
        int x = 4;

        int expected = 2;
        int actual= solution.MySqrt(x);

        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleTwo()
    {
        int x = 8;

        int expected = 2;
        int actual= solution.MySqrt(x);

        Assert.That(actual, Is.EqualTo(expected));
    }
}

