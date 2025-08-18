using Solutions.ContainerWithMostWater;

namespace Tests.ContainerWithMostWater;

public class ContainerWithMostWaterTests
{
    private ContainerWithMostWaterSolution solution;

    [SetUp]
    public void Setup() => solution = new();

    [Test]
    public void TestExampleOne()
    {
        int[] height = [1,8,6,2,5,4,8,3,7];

        int expected = 49;
        int actual = solution.MaxArea(height);
        
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleTwo()
    {
        int[] height = [1,1];

        int expected = 1;
        int actual = solution.MaxArea(height);
        
        Assert.That(actual, Is.EqualTo(expected));
    }
}

