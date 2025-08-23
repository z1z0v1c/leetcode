using Solutions.PowerOfTwo;

namespace Tests.PowerOfTwo;

public class PowerOfTwoTests
{
    private PowerOfTwoSolution solution;

    [SetUp]
    public void Setup()
    {
        solution = new PowerOfTwoSolution();
    }

    [Test]
    public void TestExampleOne()
    {
        // Arrange
        var n = 1;
        
        // Act
        var result = solution.IsPowerOfTwo(n);
        
        // Assert
        Assert.That(result, Is.True);
    }

    [Test]
    public void TestExampleTwo()
    {
        // Arrange
        var n = 16;
        
        // Act
        var result = solution.IsPowerOfTwo(n);
        
        // Assert
        Assert.That(result, Is.True);
    }

    [Test]
    public void TestExampleThree()
    {
        // Arrange
        var n = 3;
        
        // Act
        var result = solution.IsPowerOfTwo(n);
        
        // Assert
        Assert.That(result, Is.False);
    }
}