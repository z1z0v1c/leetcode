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
        const int n = 1;
        
        // Act
        var result = solution.IsPowerOfTwo(n);
        
        // Assert
        Assert.That(result, Is.True);
    }

    [Test]
    public void TestExampleTwo()
    {
        // Arrange
        const int n = 16;
        
        // Act
        var result = solution.IsPowerOfTwo(n);
        
        // Assert
        Assert.That(result, Is.True);
    }

    [Test]
    public void TestExampleThree()
    {
        // Arrange
        const int n = 3;
        
        // Act
        var result = solution.IsPowerOfTwo(n);
        
        // Assert
        Assert.That(result, Is.False);
    }

    [Test]
    public void TestExampleFour()
    {
        // Arrange
        const int n = 0;
        
        // Act
        var result = solution.IsPowerOfTwo(n);
        
        // Assert
        Assert.That(result, Is.False);
    }
}