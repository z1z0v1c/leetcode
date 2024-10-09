using Solutions.PlusOne;

namespace Tests.PlusOneTests;

public class PlusOneTests
{
    private PlusOneSolution plusOne;

    [SetUp]
    public void Setup() => plusOne = new();

    [Test]
    public void TestExampleOne()
    {
        int[] digits = [1,2,3];

        int[] expected = [1,2,4];
        int[] actual= plusOne.PlusOne(digits);

        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleTwo()
    {
        int[] digits = [4,3,2,1];

        int[] expected = [4,3,2,2];
        int[] actual= plusOne.PlusOne(digits);

        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleThree()
    {
        int[] digits = [9];

        int[] expected = [1, 0];
        int[] actual= plusOne.PlusOne(digits);

        Assert.That(actual, Is.EqualTo(expected));
    }
}
