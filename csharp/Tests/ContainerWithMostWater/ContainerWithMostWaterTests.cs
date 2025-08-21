using Solutions.ContainerWithMostWater;

namespace Tests.ContainerWithMostWater;

public class ContainerWithMostWaterTests
{
    private ContainerWithMostWaterSolution solution;

    [SetUp]
    public void Setup()
    {
        solution = new ContainerWithMostWaterSolution();
    }

    [Test]
    public void TestExampleOne()
    {
        // Arrange
        int[] height = [1, 8, 6, 2, 5, 4, 8, 3, 7];
        const int expected = 49;

        // Act
        var actual = solution.MaxArea(height);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleTwo()
    {
        // Arrange
        int[] height = [1, 1];
        const int expected = 1;

        // Act
        var actual = solution.MaxArea(height);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }
}