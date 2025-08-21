using Solutions.AddBinary;

namespace Tests.AddBinary;

public class AddBinaryTests
{
    private AddBinarySolution solution;

    [SetUp]
    public void Setup()
    {
        solution = new AddBinarySolution();
    }

    [Test]
    public void TestExampleOne()
    {
        // Arrange
        const string a = "11";
        const string b = "1";
        const string expected = "100";

        // Act
        var actual = solution.AddBinary(a, b);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleTwo()
    {
        // Arrange
        const string a = "1010";
        const string b = "1011";
        const string expected = "10101";

        // Act
        var actual = solution.AddBinary(a, b);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }
}