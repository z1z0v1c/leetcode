using Solutions.SortPackages;

namespace Tests.SortPackages;

public class SortPackagesTests
{
    private SortPackagesSolution solution;

    [SetUp]
    public void Setup()
    {
        solution = new SortPackagesSolution();
    }

    [Test]
    public void TestExampleOne()
    {
        // Arrange
        const int width = 25, height = 20, length = 20, mass = 15;
        const string expected = "STANDARD";

        // Act
        var actual = solution.Sort(width, height, length, mass);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleTwo()
    {
        // Arrange
        const int width = 149, height = 149, length = 149, mass = 15;
        const string expected = "SPECIAL";

        // Act
        var actual = solution.Sort(width, height, length, mass);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleThree()
    {
        // Arrange
        const int width = 25, height = 20, length = 151, mass = 21;
        const string expected = "REJECTED";

        // Act
        var actual = solution.Sort(width, height, length, mass);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }
}