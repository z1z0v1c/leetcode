using Solutions.IsomorphicStrings;

namespace Tests.IsomorphicStrings;

public class IsomorphicStringsTests
{
    private IsomorphicStringsSolution solution;

    [SetUp]
    public void SetUp() => solution = new();

    [Test]
    public void TestExampleOne()
    {
        // Arrange
        var s = "egg";
        var t = "add";
        
        // Act
        var result = solution.IsIsomorphic(s, t);
        
        // Assert
        Assert.That(result, Is.True);
    }
    
    [Test]
    public void TestExampleTwo()
    {
        // Arrange
        var s = "foo";
        var t = "bar";
        
        // Act
        var result = solution.IsIsomorphic(s, t);
        
        // Assert
        Assert.That(result, Is.False);
    }
    
    [Test]
    public void TestExampleThree()
    {
        // Arrange
        var s = "paper";
        var t = "title";
        
        // Act
        var result = solution.IsIsomorphic(s, t);
        
        // Assert
        Assert.That(result, Is.True);
    }
}