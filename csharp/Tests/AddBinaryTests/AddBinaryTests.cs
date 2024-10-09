using Solutions.AddBinary;

namespace Tests.AddBinaryTests;

public class AddBinaryTests
{
    private AddBinarySolution addBinary;

    [SetUp]
    public void Setup() => addBinary = new();

    [Test]
    public void TestExampleOne()
    {
        string a = "11";
        string b = "1";

        string expected = "100";
        string actual = addBinary.AddBinary(a, b);
        
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleTwo()
    {
        string a = "1010";
        string b = "1011";

        string expected = "10101";
        string actual = addBinary.AddBinary(a, b);
        
        Assert.That(actual, Is.EqualTo(expected));
    }
}
