using Solutions.SqrtX;

namespace Tests.SqrtXTests;

public class SqrtXSolutionTests
{
    private SqrtXSolution sqrtXSolution;

    [SetUp]
    public void Setup() => sqrtXSolution = new();

    [Test]
    public void TestExampleOne()
    {
        int x = 4;

        int expected = 2;
        int actual= sqrtXSolution.MySqrt(x);

        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleTwo()
    {
        int x = 8;

        int expected = 2;
        int actual= sqrtXSolution.MySqrt(x);

        Assert.That(actual, Is.EqualTo(expected));
    }
}

