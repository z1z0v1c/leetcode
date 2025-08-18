using Solutions.AddBinary;

namespace Tests.AddBinary;

public class AddBinaryTests
{
    private AddBinarySolution solution;

    [SetUp]
    public void Setup() => solution = new();

    [Test]
    public void TestExampleOne()
    {
        string a = "11";
        string b = "1";

        string expected = "100";
        string actual = solution.AddBinary(a, b);
        
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleTwo()
    {
        string a = "1010";
        string b = "1011";

        string expected = "10101";
        string actual = solution.AddBinary(a, b);
        
        Assert.That(actual, Is.EqualTo(expected));
    }
}

