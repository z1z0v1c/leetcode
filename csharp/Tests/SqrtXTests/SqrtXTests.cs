using Solutions.SqrtX;

namespace Tests.SqrtXTests;

public class SqrtXTests
{
    private SqrtX sqrtX;

    [SetUp]
    public void Setup() => sqrtX = new();

    [Test]
    public void TestExampleOne()
    {
        int x = 4;

        int expected = 2;
        int actual= sqrtX.MySqrt(x);

        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleTwo()
    {
        int x = 8;

        int expected = 2;
        int actual= sqrtX.MySqrt(x);

        Assert.That(actual, Is.EqualTo(expected));
    }
}
