using Solutions.DistributeCandies;

namespace Tests.DistributeCandies;

public class DistributeCandiesTests
{
    private DistributeCandiesSolution solution;

    [SetUp]
    public void Setup()
    {
        solution = new DistributeCandiesSolution();
    }

    [Test]
    public void TestExampleOne()
    {
        // Arrange
        const int candies = 7;
        const int numPeople = 4;
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
        const int candies = 10;
        const int numPeople = 3;
        var expected = new[] { 5, 2, 3 };

        // Act
        var actual = solution.DistributeCandies(candies, numPeople);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }
}