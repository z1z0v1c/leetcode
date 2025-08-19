using Solutions.DistributeCandies;

namespace Tests.DistributeCandies;

public class DistributeCandiesTests
{
    private DistributeCandiesSolution solution;

    [SetUp]
    public void Setup() => solution = new();

    [Test]
    public void TestExampleOne()
    {
        // Arrange
        var candies = 7;
        var numPeople = 4;
        var expected = new[] { 1, 2, 3, 1 };
        
        // Act
        var actual = solution.DistributeCandies(candies, numPeople);
        
        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }
    
    [Test]
    public void TestExampleTwo()
    {
        // Arrange
        var candies = 10;
        var numPeople = 3;
        var expected = new[] { 5, 2, 3 };
        
        // Act
        var actual = solution.DistributeCandies(candies, numPeople);
        
        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }
}