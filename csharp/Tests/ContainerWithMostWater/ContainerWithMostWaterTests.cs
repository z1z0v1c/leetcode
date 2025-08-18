using Solutions.ContainerWithMostWater;

namespace Tests.ContainerWithMostWater;

public class ContainerWithMostWaterTests
{
    private Solutions.ContainerWithMostWater.ContainerWithMostWaterSolution containerWithMostWaterSolution;

    [SetUp]
    public void Setup() => containerWithMostWaterSolution = new();

    [Test]
    public void TestExampleOne()
    {
        int[] height = [1,8,6,2,5,4,8,3,7];

        int expected = 49;
        int actual = containerWithMostWaterSolution.MaxArea(height);
        
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleTwo()
    {
        int[] height = [1,1];

        int expected = 1;
        int actual = containerWithMostWaterSolution.MaxArea(height);
        
        Assert.That(actual, Is.EqualTo(expected));
    }
}

