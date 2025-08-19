using Solutions.PeakIndexInAnArray;

namespace Tests.PeakIndexInAnArray;

public class PeakIndexInAnArrayTests
{
    private PeakIndexInAnArraySolution solution;

    [SetUp]
    public void Setup() => solution = new();

    [Test]
    public void TestExampleOne()
    {
        // Arrange
        var arr = new[] { 0, 1, 0 };
        var expected = 1;
        
        // Act
        var actual = solution.PeakIndexInMountainArray(arr);
        
        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }
    
    [Test]
    public void TestExampleTwo()
    {
        // Arrange
        var arr = new[] { 0, 2, 1, 0 };
        var expected = 1;
        
        // Act
        var actual = solution.PeakIndexInMountainArray(arr);
        
        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }
    
    [Test]
    public void TestExampleThree()
    {
        // Arrange
        var arr = new[] { 0, 10, 5, 2 };
        var expected = 1;
        
        // Act
        var actual = solution.PeakIndexInMountainArray(arr);
        
        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }
}